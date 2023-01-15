package biz

import (
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/blackhorseya/ethscan/pkg/grpcx"
	"github.com/pkg/errors"
)

func NewActivityClient(client grpcx.Client) (am.ServiceClient, error) {
	conn, err := client.Dial("activity")
	if err != nil {
		return nil, errors.Wrap(err, "activity grpc client dial error")
	}
	c := am.NewServiceClient(conn)

	return c, nil
}
