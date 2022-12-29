package repo

import (
	"time"

	bm "github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type blockRecord struct {
	Hash       string    `json:"hash" db:"hash"`
	Height     uint64    `json:"height" db:"height"`
	ParentHash string    `json:"parent_hash" db:"parent_hash"`
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`
}

func newBlockRecord(record *bm.BlockRecord) *blockRecord {
	return &blockRecord{
		Hash:       record.Hash,
		Height:     record.Height,
		ParentHash: record.ParentHash,
		Timestamp:  record.Timestamp.AsTime().UTC(),
	}
}

func (b *blockRecord) ToEntity() *bm.BlockRecord {
	return &bm.BlockRecord{
		Height:       b.Height,
		Hash:         b.Hash,
		ParentHash:   b.ParentHash,
		Transactions: nil,
		Timestamp:    timestamppb.New(b.Timestamp),
		Depth:        0,
		Status:       0,
	}
}
