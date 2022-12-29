package repo

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	columns = []string{"hash", "from", "to", "block_hash"}
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

func (s *suiteTester) Test_impl_FetchTxByHash() {
	type args struct {
		hash string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantTx  *am.Transaction
		wantErr bool
	}{
		{
			name:    "error",
			args:    args{hash: "0x38ae61626a91062204dc634319db690b48e72af453e6ff78d1866b61d41d24be"},
			wantTx:  nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTx, err := s.repo.FetchTxByHash(contextx.BackgroundWithLogger(s.logger), tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchTxByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("FetchTxByHash() gotTx = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}

func (s *suiteTester) Test_impl_CreateTx() {
	type args struct {
		tx   *am.Transaction
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create then error",
			args: args{tx: &am.Transaction{Hash: "0x0"}, mock: func() {
				s.rw.ExpectExec(`insert into txns`).WillReturnError(errors.New("error"))
			}},
			wantErr: true,
		},
		{
			name: "create then ok",
			args: args{tx: &am.Transaction{Hash: "0x0"}, mock: func() {
				s.rw.ExpectExec(`insert into txns`).WillReturnResult(sqlmock.NewResult(1, 1))
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.CreateTx(contextx.BackgroundWithLogger(s.logger), tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("CreateTx() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.assertExpectation(t)
		})
	}
}

func (s *suiteTester) Test_impl_GetTxByHash() {
	type args struct {
		hash string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantTx  *am.Transaction
		wantErr bool
	}{
		{
			name: "get tx then error",
			args: args{hash: "0x0", mock: func() {
				s.rw.ExpectQuery("select hash, `from`, `to`, block_hash from txns").
					WithArgs("0x0").WillReturnError(errors.New("error"))
			}},
			wantTx:  nil,
			wantErr: true,
		},
		{
			name: "get tx then not found",
			args: args{hash: "0x0", mock: func() {
				s.rw.ExpectQuery("select hash, `from`, `to`, block_hash from txns").
					WithArgs("0x0").
					WillReturnRows(sqlmock.NewRows([]string{"hash", "from", "to", "block_hash"}))
			}},
			wantTx:  nil,
			wantErr: false,
		},
		{
			name: "ok",
			args: args{hash: "0x0", mock: func() {
				s.rw.ExpectQuery("select hash, `from`, `to`, block_hash from txns").
					WithArgs("0x0").
					WillReturnRows(sqlmock.NewRows([]string{"hash", "from", "to", "block_hash"}).
						AddRow("0x0", "", "", ""))
			}},
			wantTx:  &am.Transaction{Hash: "0x0"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTx, err := s.repo.GetTxByHash(contextx.BackgroundWithLogger(s.logger), tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTxByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("GetTxByHash() gotTx = %v, want %v", gotTx, tt.wantTx)
			}

			s.assertExpectation(t)
		})
	}
}

func (s *suiteTester) Test_impl_ListTxns() {
	type args struct {
		cond ListTxnsCondition
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantTxns []*am.Transaction
		wantErr  bool
	}{
		{
			name: "list then error",
			args: args{cond: ListTxnsCondition{BlockHash: "0x0"}, mock: func() {
				s.rw.ExpectQuery("select hash, `from`, `to`, block_hash from txns").
					WithArgs("0x0").WillReturnError(errors.New("error"))
			}},
			wantTxns: nil,
			wantErr:  true,
		},
		{
			name: "list then not found",
			args: args{cond: ListTxnsCondition{BlockHash: "0x0"}, mock: func() {
				s.rw.ExpectQuery("select hash, `from`, `to`, block_hash from txns").
					WithArgs("0x0").WillReturnRows(sqlmock.NewRows(columns))
			}},
			wantTxns: nil,
			wantErr:  false,
		},
		{
			name: "ok",
			args: args{cond: ListTxnsCondition{BlockHash: "0x0"}, mock: func() {
				s.rw.ExpectQuery("select hash, `from`, `to`, block_hash from txns").
					WithArgs("0x0").WillReturnRows(sqlmock.NewRows(columns).AddRow(
					"",
					"",
					"",
					"0x0",
				))
			}},
			wantTxns: []*am.Transaction{{BlockHash: "0x0"}},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTxns, err := s.repo.ListTxns(contextx.BackgroundWithLogger(s.logger), tt.args.cond)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTxns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTxns, tt.wantTxns) {
				t.Errorf("ListTxns() gotTxns = %v, want %v", gotTxns, tt.wantTxns)
			}

			s.assertExpectation(t)
		})
	}
}
