package errorx

import (
	"net/http"

	"github.com/blackhorseya/ethscan/pkg/er"
)

const (
	_errServer = "Internal server error"

	_errBlockchain = "Failed to connect to blockchain"

	_errDatabase = "Failed to connect to database"

	_errKafka = "Failed to connect to kafka"
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

	// ErrFetchTx means failed to fetch transaction
	ErrFetchTx = er.New(http.StatusInternalServerError, 50012, _errBlockchain, "failed to fetch transaction")
)

var (
	// ErrGetRecord means failed to get block record
	ErrGetRecord = er.New(http.StatusInternalServerError, 50050, _errDatabase, "failed to get block record")

	// ErrCountRecord means failed to count block record
	ErrCountRecord = er.New(http.StatusInternalServerError, 50051, _errDatabase, "failed to count block record")

	// ErrCreateRecord means failed to create block record
	ErrCreateRecord = er.New(http.StatusInternalServerError, 50052, _errDatabase, "failed to create block record")
)

var (
	// ErrProduceRecord means failed to produce record to new_block
	ErrProduceRecord = er.New(http.StatusInternalServerError, 50080, _errKafka, "failed to produce record to new_block")
)
