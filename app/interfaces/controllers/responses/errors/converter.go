package errors

import (
	"fmt"
	"net/http"

	"github.com/nooboolean/seisankun_api_v2/domain"
	"github.com/nooboolean/seisankun_api_v2/domain/codes"
)

func ToHttpStatus(err error) int {
	var statusCode int
	c := domain.ErrorCode(err)

	switch c {
	case codes.BadParams, codes.EmptyBody, codes.InvalidRequest:
		statusCode = http.StatusBadRequest
	case codes.Unauthorized:
		statusCode = http.StatusUnauthorized
	case codes.NotFound:
		statusCode = http.StatusNotFound
	case codes.Forbidden:
		statusCode = http.StatusForbidden
	case codes.Database, codes.Internal, codes.Unknown:
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	switch statusCode {
	case http.StatusInternalServerError:
		fmt.Printf("stacktrace: %s\n", domain.StackTrace(err))
	}

	fmt.Printf("HttpStatus: %d, %s\n", statusCode, err)

	return statusCode
}
