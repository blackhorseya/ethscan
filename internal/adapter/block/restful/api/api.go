package api

import (
	"net/http"

	_ "github.com/blackhorseya/ethscan/api/docs" // import swagger spec
	"github.com/blackhorseya/ethscan/internal/adapter/block/restful/api/v1"
	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/blackhorseya/ethscan/pkg/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type impl struct {
}

func Handle(g *gin.RouterGroup, biz bb.IBiz) {
	g.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	i := &impl{}

	g.GET("readiness", i.Readiness)
	g.GET("liveness", i.Liveness)

	v1.Handle(g.Group("/v1"), biz)
}

func (i *impl) Readiness(c *gin.Context) {
	_, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}

func (i *impl) Liveness(c *gin.Context) {
	_, ok := c.MustGet(string(contextx.KeyCtx)).(contextx.Contextx)
	if !ok {
		_ = c.Error(errorx.ErrContextx)
		return
	}

	c.JSON(http.StatusOK, response.OK)
}
