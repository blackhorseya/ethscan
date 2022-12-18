//go:generate wire
//go:build wireinject

package repo

import (
	"github.com/blackhorseya/portto/pkg/httpx"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateRepo(opts *NodeOptions, httpclient httpx.Client, rw *sqlx.DB) IRepo {
	panic(wire.Build(testProviderSet))
}
