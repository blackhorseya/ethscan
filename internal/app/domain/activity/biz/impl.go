package biz

import (
	"github.com/blackhorseya/ethscan/internal/app/domain/activity/biz/repo"
	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(repo.ProviderSet, NewImpl)

type impl struct {
	repo repo.IRepo
}

func NewImpl(repo repo.IRepo) ab.IBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) GetByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error) {
	ret, err := i.repo.FetchTxByHash(ctx, hash)
	if err != nil {
		ctx.Error(errorx.ErrFetchTx.LogMessage, zap.Error(err), zap.String("hash", hash))
		return nil, errorx.ErrFetchTx
	}

	return ret, nil
}
