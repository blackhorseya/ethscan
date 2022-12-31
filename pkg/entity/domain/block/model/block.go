package model

import (
	"encoding/json"
)

func (x *BlockRecord) MarshalJSON() ([]byte, error) {
	type Alias BlockRecord

	var txIds []string
	for _, tx := range x.Transactions {
		txIds = append(txIds, tx.Hash)
	}

	return json.Marshal(&struct {
		*Alias
		BlockTime int64    `json:"block_time"`
		TxIds     []string `json:"transactions,omitempty"`
	}{
		Alias:     (*Alias)(x),
		BlockTime: x.Timestamp.AsTime().UTC().Unix(),
		TxIds:     txIds,
	})
}
