//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/ethscan/internal/app/domain/block/biz/repo"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo) bb.IBiz {
	panic(wire.Build(testProviderSet))
}
