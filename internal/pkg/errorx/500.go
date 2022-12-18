package errorx

import (
	"net/http"

	"github.com/blackhorseya/portto/pkg/er"
)

const (
	_errServer = "Internal server error"

	_errBlockchain = "Failed to connect to blockchain"
)

var (
	// ErrContextx means Missing contextx
	ErrContextx = er.New(http.StatusInternalServerError, 50001, _errServer, "Missing contextx")
)

var (
	// ErrFetchCurrentHeight means failed to fetch block number
	ErrFetchCurrentHeight = er.New(http.StatusInternalServerError, 50010, _errBlockchain, "failed to fetch block number")

	// ErrFetchRecord means failed to fetch block record
	ErrFetchRecord = er.New(http.StatusInternalServerError, 50011, _errBlockchain, "failed to fetch block record")
)

