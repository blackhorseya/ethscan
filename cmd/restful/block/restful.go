package main

import (
	"github.com/blackhorseya/portto/pkg/adapters"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
}

func NewRestful(logger *zap.Logger, router *gin.Engine) adapters.Restful {
	return &restful{router: router}
}

func (r *restful) InitRouting() error {
	// todo: 2022/12/15|sean|impl me

	return nil
}
