package biz

import (
	"github.com/blackhorseya/ethscan/pkg/contextx"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	bm "github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
)

// IBiz declare activity biz interface
//
//go:generate mockery --all --inpackage
type IBiz interface {
	// GetByHash serve caller to given hash to get transaction
	GetByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error)

	HandleNewBlock(ctx contextx.Contextx, record *bm.BlockRecord) (txns []*am.Transaction, err error)
}
