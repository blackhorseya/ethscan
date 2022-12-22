package app

// Service declare a service interface
//
//go:generate mockery --all --inpackage
type Service interface {
	Start() error

	AwaitSignal() error
}
