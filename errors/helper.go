package errors

import (
	"fmt"

	"github.com/morikuni/failure"
)

func wrap(err error, stackSkipCount int, wrappers ...failure.Wrapper) error {
	return failure.Custom(
		failure.Custom(err, wrappers...),
		failure.WithFormatter(),
		failure.WithCallStackSkip(stackSkipCount+1),
	)
}

func withCode(code failure.Code, err error) ServiceError {
	if err == nil {
		return nil
	}

	return &serviceError{
		code:  code,
		error: err,
	}
}

func isInvalid(code failure.Code) bool {
	_, exists := validErrors[code]
	return !exists
}

func framesToString(frames []failure.Frame) []string {
	strArr := make([]string, len(frames))
	for idx, f := range frames {
		strArr[idx] = frameToString(f)
	}

	return strArr
}

func frameToString(frame failure.Frame) string {
	return fmt.Sprintf("%+v", frame)
}

func errCodeIs(errCode failure.Code, codes ...failure.Code) bool {
	for i := range codes {
		if errCode == codes[i] {
			return true
		}
	}

	return false
}
