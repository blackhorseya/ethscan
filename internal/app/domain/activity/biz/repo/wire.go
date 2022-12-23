//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateRepo(opts *NodeOptions) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
