package blocks

import (
	bb "github.com/blackhorseya/portto/pkg/entity/domain/block/biz"
	"github.com/gin-gonic/gin"
)

type impl struct {
	biz bb.IBiz
}

func Handle(g *gin.RouterGroup, biz bb.IBiz) {
	i := &impl{biz: biz}

	g.GET(":hash", i.GetByHash)
	g.GET("", i.List)
}
