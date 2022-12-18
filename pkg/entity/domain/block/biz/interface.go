package biz

import (
	"github.com/blackhorseya/portto/pkg/contextx"
	bm "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
)

// ListCondition declare list block record condition
type ListCondition struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// IBiz declare block domain interface
//
//go:generate mockery --all --inpackage
type IBiz interface {
	// GetByHash serve caller to given hash to get block record
	GetByHash(ctx contextx.Contextx, hash string) (record *bm.BlockRecord, err error)

	// List serve caller to given condition to get block records
	List(ctx contextx.Contextx, cond ListCondition) (records []*bm.BlockRecord, err error)

	// ScanByHeight serve caller to given height to get block record
	ScanByHeight(ctx contextx.Contextx, height uint64) (record *bm.BlockRecord, next bool, err error)
}
