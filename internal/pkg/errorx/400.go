package errorx

import (
	"net/http"

	"github.com/blackhorseya/portto/pkg/er"
)

const (
	_errInvalid = "Value is invalid"
)

var (
	ErrInvalidHash = er.New(http.StatusBadRequest, 40010, _errInvalid, "cannot parse string to hash")

	ErrInvalidPage = er.New(http.StatusBadRequest, 40011, _errInvalid, "page value is invalid")

	ErrInvalidSize = er.New(http.StatusBadRequest, 40012, _errInvalid, "size value is invalid")
)
