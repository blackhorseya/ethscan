package blocks

import (
	"net/http"
	"strconv"

	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/blackhorseya/ethscan/pkg/entity/domain/block/model" // import struct
	"github.com/blackhorseya/ethscan/pkg/response"
	"github.com/gin-gonic/gin"
)

type listResponse struct {
	Total int                  `json:"total"`
	List  []*model.BlockRecord `json:"list"`
}

// List
// @Summary List block records
// @Description List block records
// @Tags Blocks
// @Accept json
// @Produce json
// @Param page query integer false "page" default(1)
// @Param size query integer false "size" default(10)
// @Success 200 {object} response.Response{data=listResponse}
// @Success 500 {object} er.Error
// @Router /v1/blocks [get]
func (i *impl) List(c *gin.Context) {
	ctx, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		_ = c.Error(err)
		return
	}

	ret, total, err := i.biz.List(ctx, bb.ListCondition{Page: page, Size: size})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.OK.WithData(listResponse{Total: total, List: ret}))
}
