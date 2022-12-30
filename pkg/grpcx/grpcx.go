package grpcx

// Server declare a grpc server functions
//
//go:generate mockery --all --inpackage
type Server interface {
	// Start a server
	Start() error

	// Stop a server
	Stop() error
}
