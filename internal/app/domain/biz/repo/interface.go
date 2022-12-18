package repo

import (
	"github.com/blackhorseya/portto/pkg/contextx"
	bm "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
	"github.com/google/wire"
)

// IRepo declare block repository interface
//
//go:generate mockery --all --inpackage
type IRepo interface {
	// FetchCurrentHeight serve caller to get current height of blockchain
	FetchCurrentHeight(ctx contextx.Contextx) (height uint64, err error)

	// FetchRecordByHeight serve caller to given height to fetch block record from node
	FetchRecordByHeight(ctx contextx.Contextx, height uint64) (record *bm.BlockRecord, err error)

	// GetRecordByHash serve caller to given hash to get block record from database
	GetRecordByHash(ctx contextx.Contextx, hash string) (record *bm.BlockRecord, err error)

	// CreateRecord serve caller to given block record to create into database
	CreateRecord(ctx contextx.Contextx, record *bm.BlockRecord) error

	// UpdateRecord serve caller to given block record to update
	UpdateRecord(ctx contextx.Contextx, record *bm.BlockRecord) error
}

var ProviderSet = wire.NewSet(NewNodeOptions, NewImpl)
