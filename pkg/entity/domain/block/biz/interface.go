package biz

import (
	"github.com/blackhorseya/ethscan/pkg/contextx"
	bm "github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
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
	List(ctx contextx.Contextx, cond ListCondition) (records []*bm.BlockRecord, total int, err error)

	// ScanBlock serve caller to scan block records
	ScanBlock(ctx contextx.Contextx, start uint64) (last uint64, progress chan *bm.BlockRecord, done chan struct{}, errC chan error)
}
