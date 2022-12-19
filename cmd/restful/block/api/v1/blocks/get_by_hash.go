package blocks

import (
	"net/http"

	"github.com/blackhorseya/portto/internal/pkg/errorx"
	"github.com/blackhorseya/portto/pkg/contextx"
	"github.com/blackhorseya/portto/pkg/response"
	"github.com/gin-gonic/gin"
)

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
// @Success 200 {object} response.Response{data=model.BlockRecord}
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

	ret, err := i.biz.GetByHash(ctx, req.Hash)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(ret))
}
