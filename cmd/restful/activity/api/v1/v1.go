package v1

import (
	"github.com/blackhorseya/ethscan/cmd/restful/activity/api/v1/transactions"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	"github.com/gin-gonic/gin"
)

func Handle(g *gin.RouterGroup, biz ab.IBiz) {
	transactions.Handle(g.Group("/transactions"), biz)
}
