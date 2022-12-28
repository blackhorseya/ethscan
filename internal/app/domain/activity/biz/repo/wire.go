//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateRepo(opts *NodeOptions, rw *sqlx.DB) (IRepo, error) {
	panic(wire.Build(testProviderSet))
}
