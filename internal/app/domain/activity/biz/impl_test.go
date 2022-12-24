package biz

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/ethscan/internal/app/domain/activity/biz/repo"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite
	logger *zap.Logger
	repo   *repo.MockIRepo
	biz    ab.IBiz
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

func (s *suiteTester) Test_impl_GetByHash() {
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
			name: "get by hash then error",
			args: args{hash: "0x0", mock: func() {
				s.repo.On("FetchTxByHash", mock.Anything, "0x0").Return(nil, errors.New("error")).Once()
			}},
			wantTx:  nil,
			wantErr: true,
		},
		{
			name: "ok",
			args: args{hash: "0x0", mock: func() {
				s.repo.On("FetchTxByHash", mock.Anything, "0x0").Return(nil, nil).Once()
			}},
			wantTx:  nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotTx, err := s.biz.GetByHash(contextx.BackgroundWithLogger(s.logger), tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("GetByHash() gotTx = %v, want %v", gotTx, tt.wantTx)
			}

			s.assertExpectation(t)
		})
	}
}
