package grpcx

import (
	"google.golang.org/grpc"
)

// Server declare a grpc server functions
//
//go:generate mockery --all --inpackage
type Server interface {
	// Start a server
	Start() error

	// Stop a server
	Stop() error
}

// Client declare a grpc client functions
//
//go:generate mockery --all --inpackage
type Client interface {
	Dial(service string) (*grpc.ClientConn, error)
}
