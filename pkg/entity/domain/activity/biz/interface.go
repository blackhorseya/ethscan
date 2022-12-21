package biz

import (
	"github.com/blackhorseya/portto/pkg/contextx"
	am "github.com/blackhorseya/portto/pkg/entity/domain/activity/model"
)

// IBiz declare activity biz interface
//
//go:generate mockery --all --inpackage
type IBiz interface {
	// GetByHash serve caller to given hash to get transaction
	GetByHash(ctx contextx.Contextx, hash string) (tx *am.Transaction, err error)
}
