package repo

import (
	"github.com/google/wire"
)

// IRepo declare activity repository interface
//go:generate mockery --all --inpackage
type IRepo interface {

}

var ProviderSet = wire.NewSet()
