package model

import (
	"encoding/json"
)

func (x *BlockRecord) MarshalJSON() ([]byte, error) {
	type Alias BlockRecord

	return json.Marshal(&struct {
		*Alias
		BlockTime int64 `json:"block_time"`
	}{
		Alias:     (*Alias)(x),
		BlockTime: x.Timestamp.AsTime().UTC().Unix(),
	})
}
