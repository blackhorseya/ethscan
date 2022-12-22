package repo

import (
	"github.com/blackhorseya/portto/pkg/contextx"
	bm "github.com/blackhorseya/portto/pkg/entity/domain/block/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/wire"
)

type ListRecordCondition struct {
	Limit  int
	Offset int
}

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

	// ListRecord serve caller to given condition to list block record from database
	ListRecord(ctx contextx.Contextx, condition ListRecordCondition) (records []*bm.BlockRecord, err error)

	// CountRecord serve caller to given condition to count block records from database
	CountRecord(ctx contextx.Contextx, condition ListRecordCondition) (total int, err error)

	// CreateRecord serve caller to given block record to create into database
	CreateRecord(ctx contextx.Contextx, record *bm.BlockRecord) error

	// PublishRecord serve caller to given record to publish the record to new_block
	PublishRecord(ctx contextx.Contextx, record *bm.BlockRecord, delivery chan kafka.Event) error
}

var ProviderSet = wire.NewSet(NewNodeOptions, NewImpl)
