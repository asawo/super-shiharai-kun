package errors

import (
	"errors"
	"fmt"
	"strings"

	"github.com/morikuni/failure"
)

type ServiceError interface {
	GetCode() failure.Code
	GetError() error

	ErrorCodeIs(codes ...failure.Code) bool

	Kind() string
	Message() string
	StackTrace() []string

	Error() string
}

type serviceError struct {
	code  failure.Code
	error error
}

func New(code failure.Code, wrappers ...failure.Wrapper) ServiceError {
	emptyErr := errors.New("")
	wrappedErr := wrap(emptyErr, 1, wrappers...)

	if isInvalid(code) {
		code = Unknown
	}

	return withCode(code, wrappedErr)
}

func NewFromError(err error, code failure.Code, wrappers ...failure.Wrapper) ServiceError {
	wrappedErr := wrap(err, 1, wrappers...)

	if isInvalid(code) {
		code = Unknown
	}

	return withCode(code, wrappedErr)
}

func Wrap(se ServiceError, wrappers ...failure.Wrapper) ServiceError {
	wrappedErr := wrap(se.GetError(), 1, wrappers...)
	return withCode(se.GetCode(), wrappedErr)
}

func (se *serviceError) GetCode() failure.Code {
	return se.code
}

func (se *serviceError) GetError() error {
	return se.error
}

func (se *serviceError) Kind() string {
	return se.GetCode().ErrorCode()
}

func (se *serviceError) Message() string {
	errString := se.errString()
	return errString
}

func (se *serviceError) StackTrace() []string {
	cs, _ := failure.CallStackOf(se.error)
	return framesToString(cs.Frames())
}

func (se *serviceError) Error() string {
	errString := se.errString()
	return fmt.Sprintf("code(%s): %s", se.code, errString)
}

func (se *serviceError) errString() string {
	errString := se.error.Error()
	errString = strings.TrimSpace(errString)
	errString = strings.TrimSuffix(errString, ":")
	return errString
}

func (se *serviceError) ErrorCodeIs(codes ...failure.Code) bool {
	return errCodeIs(se.GetCode(), codes...)
}
