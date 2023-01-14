package repo

import (
	"encoding/json"
	"time"

	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
)

type transaction struct {
	Hash      string    `json:"hash" db:"hash"`
	From      string    `json:"from" db:"from"`
	To        string    `json:"to" db:"to"`
	BlockHash string    `json:"block_hash" db:"block_hash"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Nonce     uint64    `json:"nonce" db:"nonce"`
	Data      string    `json:"data" db:"data"`
	Value     string    `json:"value" db:"value"`
	Events    string    `json:"events" db:"events"`
}

func newTransaction(val *am.Transaction) (*transaction, error) {
	events, err := json.Marshal(val.Events)
	if err != nil {
		return nil, err
	}

	return &transaction{
		Hash:      val.Hash,
		From:      val.From,
		To:        val.To,
		BlockHash: val.BlockHash,
		Timestamp: time.Time{},
		Nonce:     val.Nonce,
		Data:      val.Data,
		Value:     val.Value,
		Events:    string(events),
	}, nil
}

func (t *transaction) ToEntity() *am.Transaction {
	var events []*am.Event
	_ = json.Unmarshal([]byte(t.Events), &events)

	return &am.Transaction{
		BlockHash: t.BlockHash,
		Hash:      t.Hash,
		From:      t.From,
		To:        t.To,
		Nonce:     t.Nonce,
		Data:      t.Data,
		Value:     t.Value,
		Events:    events,
	}
}
