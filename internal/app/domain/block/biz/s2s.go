package biz

import (
	"github.com/blackhorseya/ethscan/pkg/entity/domain/activity/s2s"
	"github.com/blackhorseya/ethscan/pkg/grpcx"
	"github.com/pkg/errors"
)

func NewActivityClient(client grpcx.Client) (s2s.ServiceClient, error) {
	conn, err := client.Dial("activity")
	if err != nil {
		return nil, errors.Wrap(err, "activity grpc client dial error")
	}
	c := s2s.NewServiceClient(conn)

	return c, nil
}
