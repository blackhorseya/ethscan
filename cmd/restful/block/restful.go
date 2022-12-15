package main

import (
	"github.com/blackhorseya/portto/cmd/restful/block/api"
	"github.com/blackhorseya/portto/pkg/adapters"
	"github.com/blackhorseya/portto/pkg/contextx"
	"github.com/blackhorseya/portto/pkg/er"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
}

func NewRestful(logger *zap.Logger, router *gin.Engine) adapters.Restful {
	router.Use(contextx.AddContextxWitLoggerMiddleware(logger))
	router.Use(er.AddErrorHandlingMiddleware())

	return &restful{router: router}
}

func (r *restful) InitRouting() error {
	api.Handle(r.router.Group("/api"))

	return nil
}
