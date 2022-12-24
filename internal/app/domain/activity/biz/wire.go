//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/ethscan/internal/app/domain/activity/biz/repo"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo) ab.IBiz {
	panic(wire.Build(testProviderSet))
}
