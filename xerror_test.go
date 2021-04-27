package xerror

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

type testHandler struct{}

func TestXerror(t *testing.T) {
	// register
	Register(&testHandler{})

	if testThrowError() == nil {
		t.Fail()
	}

	if testThrowErrorWhen() == nil {
		t.Fail()
	}

	if testThrowf() == nil {
		t.Fail()
	}

	if testThrowWhen() == nil {
		t.Fail()
	}

}

func testThrowError() (err error) {
	defer AutoHandle(context.Background(), &err)
	e := errors.New("TestThrowError error")

	ThrowError(e)
	return nil
}

func testThrowErrorWhen() (err error) {
	defer AutoHandle(context.Background(), &err)
	e := errors.New("TestThrowError error")

	ThrowErrorWhen(1 > 0, e)
	return nil
}

func testThrowf() (err error) {
	defer AutoHandle(context.Background(), &err)

	Throwf("format key: %s", "key")
	return nil
}

func testThrowWhen() (err error) {
	defer AutoHandle(context.Background(), &err)

	ThrowWhen(1 > 0, "format key: %s", "key")
	return nil
}

func (t testHandler) Match(ctx context.Context, i interface{}) (match bool) {
	return true
}

func (t testHandler) Handle(ctx context.Context, i interface{}) (stop bool) {
	return false
}

func (t testHandler) ConvertToError(ctx context.Context, i interface{}) error {
	return errors.New(fmt.Sprintf("%v", i))
}
