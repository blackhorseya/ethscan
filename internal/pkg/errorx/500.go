package errorx

import (
	"net/http"

	"github.com/blackhorseya/portto/pkg/er"
)

const (
	_errServer = "Internal server error"
)

var (
	// ErrContextx means Missing contextx
	ErrContextx = er.New(http.StatusInternalServerError, 50001, _errServer, "Missing contextx")
)
