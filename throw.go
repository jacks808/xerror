package xerror

import (
	"errors"
	"fmt"
)

// throw xerror with format message
func Throwf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	panic(newXerror(msg))
}

// throw a Xerror when condition is true
func ThrowWhen(condition bool, format string, args ...interface{}) {
	if condition {
		Throwf(format, args...)
	}
}

// throw any error when condition is true
func ThrowErrorWhen(condition bool, err error) {
	if condition {
		ThrowError(err)
	}
}

// throw any error when err is not nil
func ThrowError(err error) {
	if err == nil {
		return
	}

	if e, ok := err.(*Xerror); ok {
		Throwf("%s", e)
	}

	panic(err)
}

// convert string msg to Xerror
func newXerror(msg string) error {
	e := &Xerror{
		error: errors.New(msg),
	}
	return e
}
