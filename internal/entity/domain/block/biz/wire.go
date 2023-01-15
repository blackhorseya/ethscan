//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/ethscan/internal/entity/domain/block/biz/repo"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo, activity am.ServiceClient) bb.IBiz {
	panic(wire.Build(testProviderSet))
}
