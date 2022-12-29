package repo

import (
	"github.com/blackhorseya/ethscan/pkg/contextx"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/google/wire"
)

type ListTxnsCondition struct {
	BlockHash string `json:"block_hash"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// IRepo declare activity repository interface
//
//go:generate mockery --all --inpackage
type IRepo interface {
	// FetchTxByHash serve caller to given hash to get transaction from rpc
	FetchTxByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error)

	CreateTx(ctx contextx.Contextx, tx *am.Transaction) error

	// GetTxByHash serve caller to given hash to get transaction
	GetTxByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error)

	// ListTxns serve caller to given condition to list txns
	ListTxns(ctx contextx.Contextx, cond ListTxnsCondition) (txns []*am.Transaction, err error)
}

var ProviderSet = wire.NewSet(NewNodeOptions, NewImpl)
