package repo

import (
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
)

type transaction struct {
	Hash      string `json:"hash" db:"hash"`
	From      string `json:"from" db:"from"`
	To        string `json:"to" db:"to"`
	BlockHash string `json:"block_hash" db:"block_hash"`
}

func newTransaction(val *am.Transaction) *transaction {
	return &transaction{
		Hash:      val.Hash,
		From:      val.From,
		To:        val.To,
		BlockHash: val.BlockHash,
	}
}

func (t *transaction) ToEntity() *am.Transaction {
	return &am.Transaction{
		BlockHash: t.BlockHash,
		Hash:      t.Hash,
		From:      t.From,
		To:        t.To,
		Nonce:     0,
		Data:      "",
		Value:     "",
		Events:    nil,
	}
}
