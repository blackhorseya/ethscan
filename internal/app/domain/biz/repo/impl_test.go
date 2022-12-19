package repo

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/portto/pkg/contextx"
	bm "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	columns = []string{"hash", "height", "parent_hash", "timestamp"}

	now = time.Now()
)

type suiteTester struct {
	suite.Suite
	logger *zap.Logger
	rw     sqlmock.Sqlmock
	repo   IRepo
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	db, rw, _ := sqlmock.New(sqlmock.MonitorPingsOption(true))
	s.rw = rw

	opts := &NodeOptions{BaseURL: "http://localhost", Timeout: 5 * time.Second}
	s.repo, _ = CreateRepo(opts, sqlx.NewDb(db, "mysql"))
}

func (s *suiteTester) assertExpectation(t *testing.T) {
	if err := s.rw.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetRecordByHash() {
	type args struct {
		hash string
		mock func()
	}
	tests := []struct {
		name       string
		args       args
		wantRecord *bm.BlockRecord
		wantErr    bool
	}{
		{
			name: "get record by hash then error",
			args: args{hash: "0x0", mock: func() {
				s.rw.ExpectQuery(`select hash, height, parent_hash, timestamp from records`).
					WithArgs("0x0").
					WillReturnError(errors.New("error"))
			}},
			wantRecord: nil,
			wantErr:    true,
		},
		{
			name: "get record by hash then not found",
			args: args{hash: "0x0", mock: func() {
				s.rw.ExpectQuery(`select hash, height, parent_hash, timestamp from records`).
					WithArgs("0x0").
					WillReturnRows(sqlmock.NewRows(columns))
			}},
			wantRecord: nil,
			wantErr:    false,
		},
		{
			name: "ok",
			args: args{hash: "0x0", mock: func() {
				s.rw.ExpectQuery(`select hash, height, parent_hash, timestamp from records`).
					WithArgs("0x0").
					WillReturnRows(sqlmock.NewRows(columns).AddRow(
						"hash",
						uint64(0),
						"parent",
						now,
					))
			}},
			wantRecord: &bm.BlockRecord{
				Height:         0,
				Hash:           "hash",
				ParentHash:     "parent",
				TransactionIds: nil,
				Timestamp:      timestamppb.New(now),
				Depth:          0,
				Status:         0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecord, err := s.repo.GetRecordByHash(contextx.BackgroundWithLogger(s.logger), tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecordByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("GetRecordByHash() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}

			s.assertExpectation(t)
		})
	}
}

func (s *suiteTester) Test_impl_CreateRecord() {
	type args struct {
		record *bm.BlockRecord
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create record then error",
			args: args{record: &bm.BlockRecord{Hash: "hash", Height: 0, ParentHash: "parent", Timestamp: timestamppb.New(now)}, mock: func() {
				s.rw.ExpectExec(`insert into records`).
					WithArgs("hash", uint64(0), "parent", sqlmock.AnyArg()).
					WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "ok",
			args: args{record: &bm.BlockRecord{Hash: "hash", Height: 0, ParentHash: "parent", Timestamp: timestamppb.New(now)}, mock: func() {
				s.rw.ExpectExec(`insert into records`).
					WithArgs("hash", uint64(0), "parent", sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.CreateRecord(contextx.BackgroundWithLogger(s.logger), tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("CreateRecord() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.assertExpectation(t)
		})
	}
}
