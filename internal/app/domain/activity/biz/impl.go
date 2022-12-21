package biz

import (
	"github.com/blackhorseya/portto/internal/app/domain/activity/biz/repo"
	"github.com/blackhorseya/portto/pkg/contextx"
	ab "github.com/blackhorseya/portto/pkg/entity/domain/activity/biz"
	"github.com/blackhorseya/portto/pkg/entity/domain/activity/model"
	"github.com/google/wire"
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

func (i *impl) GetByHash(ctx contextx.Contextx, hash string) (tx *model.Transaction, err error) {
	// todo: 2022/12/21|sean|impl me
	panic("implement me")
}
