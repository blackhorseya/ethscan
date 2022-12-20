package main

import (
	"time"

	"github.com/blackhorseya/portto/pkg/adapters"
	"github.com/blackhorseya/portto/pkg/contextx"
	"github.com/blackhorseya/portto/pkg/er"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
}

func NewRestful(logger *zap.Logger, router *gin.Engine) adapters.Restful {
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/api/readiness", "/api/liveness"},
	}))

	router.Use(contextx.AddContextxWitLoggerMiddleware(logger))
	router.Use(er.AddErrorHandlingMiddleware())

	return &restful{router: router}
}

func (r *restful) InitRouting() error {
	// todo: 2022/12/20|sean|impl me
	return nil
}
