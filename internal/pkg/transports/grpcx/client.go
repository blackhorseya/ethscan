package grpcx

import (
	"fmt"

	"github.com/blackhorseya/ethscan/pkg/contextx"
	"github.com/blackhorseya/ethscan/pkg/grpcx"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientOptions struct {
	Service map[string]string `json:"service" yaml:"service"`
}

func NewClientOptions(v *viper.Viper) (*ClientOptions, error) {
	opts := new(ClientOptions)

	err := v.UnmarshalKey("grpc", &opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

type client struct {
	opts *ClientOptions
}

func NewClient(opts *ClientOptions) grpcx.Client {
	return &client{
		opts: opts,
	}
}

func (c *client) Dial(service string) (*grpc.ClientConn, error) {
	ctx := contextx.Background()

	target, ok := c.opts.Service[service]
	if !ok {
		return nil, fmt.Errorf("cannot find service name: %s", service)
	}

	conn, err := grpc.DialContext(ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "grpc client dial error")
	}

	return conn, nil
}

var ProviderClient = wire.NewSet(NewClientOptions, NewClient)
