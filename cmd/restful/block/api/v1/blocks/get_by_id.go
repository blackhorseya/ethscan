package blocks

import (
	"github.com/gin-gonic/gin"
)

// GetByHash
// @Summary Get a block record by hash
// @Description Get a block record by hash
// @Tags Blocks
// @Accept json
// @Produce json
// @Param hash path string true "hash"
// @Success 200 {object} response.Response{data=model.BlockRecord}
// @Success 500 {object} er.Error
// @Router /v1/blocks/{hash} [get]
func (i *impl) GetByHash(c *gin.Context) {
	// todo: 2022/12/19|sean|impl me
}
