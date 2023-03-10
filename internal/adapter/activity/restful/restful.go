package restful

import (
	"time"

	"github.com/blackhorseya/ethscan/internal/adapter/activity/restful/api"
	"github.com/blackhorseya/ethscan/pkg/adapters"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	"github.com/blackhorseya/ethscan/pkg/er"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
	biz    ab.IBiz
}

func NewRestful(logger *zap.Logger, router *gin.Engine, biz ab.IBiz) adapters.Restful {
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/readiness", "/api/liveness"},
	}))

	router.Use(contextx.AddContextxWitLoggerMiddleware(logger))
	router.Use(er.AddErrorHandlingMiddleware())

	return &restful{router: router, biz: biz}
}

func (r *restful) InitRouting() error {
	api.Handle(r.router.Group("/api"), r.biz)

	return nil
}

var ActivitySet = wire.NewSet(NewRestful)
