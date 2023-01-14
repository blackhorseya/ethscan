package v1

import (
	"github.com/blackhorseya/ethscan/internal/adapter/block/restful/api/v1/blocks"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz bb.IBiz) {
	blocks.Handle(g.Group("/blocks"), biz)
}
