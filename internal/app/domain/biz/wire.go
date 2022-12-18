//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/portto/internal/app/domain/biz/repo"
	bb "github.com/blackhorseya/portto/pkg/entity/domain/block/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo) bb.IBiz {
	panic(wire.Build(testProviderSet))
}
