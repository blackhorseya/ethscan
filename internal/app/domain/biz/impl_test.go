package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/portto/internal/app/domain/biz/repo"
	"github.com/blackhorseya/portto/pkg/contextx"
	bb "github.com/blackhorseya/portto/pkg/entity/domain/block/biz"
	bm "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite
	logger *zap.Logger
	repo   *repo.MockIRepo
	biz    bb.IBiz
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.repo = new(repo.MockIRepo)
	s.biz = CreateBiz(s.repo)
}

func (s *suiteTester) assertExpectation(t *testing.T) {
	s.repo.AssertExpectations(t)
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_ScanByHeight() {
	type args struct {
		height uint64
		mock   func()
	}
	tests := []struct {
		name       string
		args       args
		wantRecord *bm.BlockRecord
		wantNext   bool
		wantErr    bool
	}{
		{
			name: "fetch current height then error",
			args: args{height: 0, mock: func() {
				s.repo.On("FetchCurrentHeight", mock.Anything).Return(uint64(0), errors.New("error")).Once()
			}},
			wantRecord: nil,
			wantNext:   false,
			wantErr:    true,
		},
		{
			name: "fetch block by height then error",
			args: args{height: 0, mock: func() {
				s.repo.On("FetchCurrentHeight", mock.Anything).Return(uint64(100), nil).Once()

				s.repo.On("FetchRecordByHeight", mock.Anything, uint64(0)).Return(nil, errors.New("error")).Once()
			}},
			wantRecord: nil,
			wantNext:   false,
			wantErr:    true,
		},
		{
			name: "ok",
			args: args{height: 0, mock: func() {
				s.repo.On("FetchCurrentHeight", mock.Anything).Return(uint64(100), nil).Once()

				s.repo.On("FetchRecordByHeight", mock.Anything, uint64(0)).Return(&bm.BlockRecord{Height: 0}, nil).Once()
			}},
			wantRecord: &bm.BlockRecord{Height: 0},
			wantNext:   true,
			wantErr:    false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotRecord, gotNext, err := s.biz.ScanByHeight(contextx.BackgroundWithLogger(s.logger), tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScanByHeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecord, tt.wantRecord) {
				t.Errorf("ScanByHeight() gotRecord = %v, want %v", gotRecord, tt.wantRecord)
			}
			if gotNext != tt.wantNext {
				t.Errorf("ScanByHeight() gotNext = %v, want %v", gotNext, tt.wantNext)
			}

			s.assertExpectation(t)
		})
	}
}
