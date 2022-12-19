package v1

import (
	"github.com/blackhorseya/portto/cmd/restful/block/api/v1/blocks"
	bb "github.com/blackhorseya/portto/pkg/entity/domain/block/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz bb.IBiz) {
	blocks.Handle(g.Group("/blocks"), biz)
}
