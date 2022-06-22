package domain

import (
	"fmt"

	"github.com/nooboolean/seisankun_api_v2/domain/codes"
	"github.com/pkg/errors"
)

type privateError struct {
	errorCode codes.ErrorCode
	err       error
}

func (e privateError) Error() string {
	return fmt.Sprintf("ErrorCode: %s, Msg: %s", e.errorCode, e.err)
}

func Errorf(c codes.ErrorCode, format string, a ...interface{}) error {
	if c == codes.OK {
		return nil
	}
	return privateError{
		errorCode: c,
		err:       errors.Errorf(format, a...),
	}
}

func ErrorCode(err error) codes.ErrorCode {
	if err == nil {
		return codes.OK
	}
	var e privateError
	if errors.As(err, &e) {
		return e.errorCode
	}
	return codes.Unknown
}

func StackTrace(err error) string {
	var e privateError
	if errors.As(err, &e) {
		return fmt.Sprintf("%+v\n", e.err)
	}
	return ""
}
