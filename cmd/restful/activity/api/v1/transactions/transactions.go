package transactions

import (
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ab.IBiz) {
	i := &impl{biz: biz}

	g.GET("/:hash", i.GetByHash)
}

type impl struct {
	biz ab.IBiz
}
