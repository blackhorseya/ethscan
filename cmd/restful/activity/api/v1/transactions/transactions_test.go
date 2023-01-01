package transactions

import (
	"testing"

	"github.com/blackhorseya/ethscan/pkg/contextx"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	"github.com/blackhorseya/ethscan/pkg/er"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite
	logger *zap.Logger
	r      *gin.Engine
	biz    *ab.MockIBiz
	handle *impl
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.biz = new(ab.MockIBiz)

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(contextx.AddContextxWitLoggerMiddleware(s.logger))
	s.r.Use(er.AddErrorHandlingMiddleware())
	Handle(s.r.Group("/api/v1/transactions"), s.biz)
}

func (s *suiteTester) assertExpectation(t *testing.T) {
	s.biz.AssertExpectations(t)
}

func TestAll(t *testing.T) {
	suite.Run(t, new(suiteTester))
}
