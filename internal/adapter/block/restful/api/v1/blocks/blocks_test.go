package blocks

import (
	"testing"

	"github.com/blackhorseya/ethscan/pkg/contextx"
	bb "github.com/blackhorseya/ethscan/pkg/entity/domain/block/biz"
	"github.com/blackhorseya/ethscan/pkg/er"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite
	logger *zap.Logger
	r      *gin.Engine
	biz    *bb.MockIBiz
	handle *impl
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.biz = new(bb.MockIBiz)

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(contextx.AddContextxWitLoggerMiddleware(s.logger))
	s.r.Use(er.AddErrorHandlingMiddleware())
	Handle(s.r.Group("/api/v1/blocks"), s.biz)
}

func (s *suiteTester) assertExpectation(t *testing.T) {
	s.biz.AssertExpectations(t)
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
