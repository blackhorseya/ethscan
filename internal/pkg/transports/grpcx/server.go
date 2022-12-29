package grpcx

import (
	"fmt"
	"net"

	"github.com/blackhorseya/ethscan/pkg/contextx"
	"github.com/blackhorseya/ethscan/pkg/grpcx"
	"github.com/blackhorseya/ethscan/pkg/netx"
	"github.com/google/wire"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// todo: 2022/12/30|sean|breakpoint: implement grpc grpcserver return grpx grpcserver

type ServerOptions struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
}

func NewServerOptions(v *viper.Viper) (*ServerOptions, error) {
	opts := new(ServerOptions)

	err := v.UnmarshalKey("grpc", &opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

type server struct {
	logger *zap.Logger

	host       string
	port       int
	grpcserver *grpc.Server
}

func NewServer(opts *ServerOptions, logger *zap.Logger) grpcx.Server {
	stream := grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
			contextx.StreamServerInterceptor(logger),
		),
	)
	unary := grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
			contextx.UnaryServerInterceptor(logger),
		),
	)
	gs := grpc.NewServer(stream, unary)

	return &server{
		logger:     logger,
		host:       opts.Host,
		port:       opts.Port,
		grpcserver: gs,
	}
}

func (s *server) Start() error {
	if s.host == "" {
		s.host = "0.0.0.0"
	}

	if s.port == 0 {
		s.port = netx.GetAvailablePort()
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.logger.Info("grpc server starting...", zap.String("addr", addr))

	go func() {
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			s.logger.Fatal("net listen", zap.Error(err), zap.String("addr", addr))
			return
		}

		err = s.grpcserver.Serve(listen)
		if err != nil {
			s.logger.Fatal("start grpc server error", zap.Error(err))
			return
		}
	}()

	return nil
}

func (s *server) Stop() error {
	s.logger.Info("grpc server stopping...")

	s.grpcserver.GracefulStop()

	return nil
}

var ProviderServer = wire.NewSet(NewServerOptions, NewServer)
