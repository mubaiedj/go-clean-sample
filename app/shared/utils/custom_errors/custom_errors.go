package custom_errors

import (
	"errors"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/metrics"
)

const (
	DataBaseError = "DATA_BASE_ERROR"
	Unknown       = "UNKNOWN"
)

type RequestError struct {
	kind string
	err  error
}

func New(message string, kind string) error {
	log.Error("[%s] %s", kind, message)
	metrics.IncrementErrors(kind)
	return &RequestError{
		kind: kind,
		err:  errors.New(message),
	}
}

func NewWithError(err error, kind string) error {
	log.Error("[%s] %s", kind, err.Error())
	metrics.IncrementErrors(kind)
	return &RequestError{
		kind: kind,
		err:  err,
	}
}

func (custom *RequestError) Error() string {
	return custom.err.Error()
}

func (custom *RequestError) Kind() string {
	return custom.kind
}
