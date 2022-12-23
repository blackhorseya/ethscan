package repo

import (
	"github.com/blackhorseya/ethscan/pkg/contextx"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/google/wire"
)

// IRepo declare activity repository interface
//
//go:generate mockery --all --inpackage
type IRepo interface {
	// FetchTxByHash serve caller to given hash to get transaction from rpc
	FetchTxByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error)
}

var ProviderSet = wire.NewSet(NewNodeOptions, NewImpl)
