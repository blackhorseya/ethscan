package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/ethscan/internal/entity/domain/block/biz/repo"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	"github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	bm "github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite
	logger   *zap.Logger
	repo     *repo.MockIRepo
	activity *model.MockServiceClient
	biz      bb.IBiz
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.activity = new(model.MockServiceClient)
	s.repo = new(repo.MockIRepo)
	s.biz = CreateBiz(s.repo, s.activity)
}

func (s *suiteTester) assertExpectation(t *testing.T) {
	s.repo.AssertExpectations(t)
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetByHash() {
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
			args: args{hash: "hash", mock: func() {
				s.repo.On("GetRecordByHash", mock.Anything, "hash").Return(nil, errors.New("error")).Once()
			}},
			wantRecord: nil,
			wantErr:    true,
		},
		{
			name: "list txns by block hash then ok",
			args: args{hash: "hash", mock: func() {
				s.repo.On("GetRecordByHash", mock.Anything, "hash").Return(&bm.BlockRecord{Hash: "hash"}, nil).Once()

				s.activity.On("ListTxnsByBlockHash", mock.Anything, &model.ListTxnsByBlockHashRequest{Hash: "hash"}).Return(nil, errors.New("error")).Once()
			}},
			wantRecord: &bm.BlockRecord{Hash: "hash"},
			wantErr:    false,
		},
		{
			name: "ok",
			args: args{hash: "hash", mock: func() {
				s.repo.On("GetRecordByHash", mock.Anything, "hash").Return(&bm.BlockRecord{Hash: "hash"}, nil).Once()

				s.activity.On("ListTxnsByBlockHash", mock.Anything, &model.ListTxnsByBlockHashRequest{Hash: "hash"}).
					Return(&model.ListTxnsByBlockHashResponse{}, nil).Once()
			}},
			wantRecord: &bm.BlockRecord{Hash: "hash"},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecord, err := s.biz.GetByHash(contextx.BackgroundWithLogger(s.logger), tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("GetByHash() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}

			s.assertExpectation(t)
		})
	}
}

func (s *suiteTester) Test_impl_List() {
	type args struct {
		cond bb.ListCondition
		mock func()
	}
	tests := []struct {
		name        string
		args        args
		wantRecords []*bm.BlockRecord
		wantTotal   int
		wantErr     bool
	}{
		{
			name:        "invalid page",
			args:        args{cond: bb.ListCondition{Page: -1, Size: 10}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name:        "invalid size",
			args:        args{cond: bb.ListCondition{Page: 1, Size: -1}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "list then error",
			args: args{cond: bb.ListCondition{Page: 1, Size: 10}, mock: func() {
				s.repo.On("ListRecord", mock.Anything, repo.ListRecordCondition{Limit: 10, Offset: 0}).Return(nil, errors.New("error")).Once()
			}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "count then error",
			args: args{cond: bb.ListCondition{Page: 1, Size: 10}, mock: func() {
				s.repo.On("ListRecord", mock.Anything, repo.ListRecordCondition{Limit: 10, Offset: 0}).Return(nil, nil).Once()

				s.repo.On("CountRecord", mock.Anything, repo.ListRecordCondition{Limit: 10, Offset: 0}).Return(0, errors.New("error")).Once()
			}},
			wantRecords: nil,
			wantTotal:   0,
			wantErr:     true,
		},
		{
			name: "ok",
			args: args{cond: bb.ListCondition{Page: 1, Size: 10}, mock: func() {
				s.repo.On("ListRecord", mock.Anything, repo.ListRecordCondition{Limit: 10, Offset: 0}).Return(nil, nil).Once()

				s.repo.On("CountRecord", mock.Anything, repo.ListRecordCondition{Limit: 10, Offset: 0}).Return(10, nil).Once()
			}},
			wantRecords: nil,
			wantTotal:   10,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecords, gotTotal, err := s.biz.List(contextx.BackgroundWithLogger(s.logger), tt.args.cond)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("List() gotRecords = %v, want %v", gotRecords, tt.wantRecords)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("List() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}

			s.assertExpectation(t)
		})
	}
}
