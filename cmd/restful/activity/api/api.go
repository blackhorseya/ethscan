package api

import (
	_ "github.com/blackhorseya/portto/api/docs" // import swagger spec
	ab "github.com/blackhorseya/portto/pkg/entity/domain/activity/biz"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Handle(g *gin.RouterGroup, biz ab.IBiz) {
	if gin.Mode() != gin.ReleaseMode {
		g.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
