package api

import (
	"net/http"

	_ "github.com/blackhorseya/portto/api/docs" // import swagger spec
	"github.com/blackhorseya/portto/internal/pkg/errorx"
	"github.com/blackhorseya/portto/pkg/contextx"
	"github.com/blackhorseya/portto/pkg/response"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type impl struct {
}

func Handle(g *gin.RouterGroup) {
	i := &impl{}

	if gin.Mode() != gin.ReleaseMode {
		g.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	g.GET("readiness", i.Readiness)
	g.GET("liveness", i.Liveness)
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
