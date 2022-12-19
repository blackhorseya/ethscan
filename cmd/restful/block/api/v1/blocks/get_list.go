package blocks

import (
	"github.com/blackhorseya/portto/pkg/entity/domain/block/model" // import struct
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
	// todo: 2022/12/19|sean|impl me
}
