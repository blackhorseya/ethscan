package blocks

import (
	"encoding/json"
	"net/http"

	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	"github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
	"github.com/blackhorseya/ethscan/pkg/response"
	"github.com/gin-gonic/gin"
)

type blockResponse struct {
	Block *model.BlockRecord
}

func (x *blockResponse) MarshalJSON() ([]byte, error) {
	type Alias model.BlockRecord

	var txns []string
	if x.Block != nil {
		for _, tx := range x.Block.Transactions {
			txns = append(txns, tx.Hash)
		}
	}

	var timestamp int64
	if x.Block != nil {
		timestamp = x.Block.Timestamp.AsTime().UTC().Unix()
	}

	return json.Marshal(&struct {
		*Alias
		BlockTime    int64    `json:"block_time"`
		Transactions []string `json:"transactions,omitempty"`
	}{
		Alias:        (*Alias)(x.Block),
		BlockTime:    timestamp,
		Transactions: txns,
	})
}

type getByHashRequest struct {
	Hash string `uri:"hash" binding:"required"`
}

// GetByHash
// @Summary Get a block record by hash
// @Description Get a block record by hash
// @Tags Blocks
// @Accept json
// @Produce json
// @Param hash path string true "hash"
// @Success 200 {object} response.Response{data=blockResponse}
// @Success 500 {object} er.Error
// @Router /v1/blocks/{hash} [get]
func (i *impl) GetByHash(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	var req *getByHashRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		_ = c.Error(err)
		return
	}

	block, err := i.biz.GetByHash(ctx, req.Hash)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(&blockResponse{Block: block}))
}
