//go:generate wire
//go:build wireinject

package biz

import (
	"github.com/blackhorseya/ethscan/internal/entity/domain/block/biz/repo"
	"github.com/blackhorseya/ethscan/pkg/entity/domain/activity/s2s"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IRepo, activity s2s.ServiceClient) bb.IBiz {
	panic(wire.Build(testProviderSet))
}
