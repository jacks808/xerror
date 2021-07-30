package xerror

import (
	"context"
	"errors"
	"fmt"
	"testing"
)


// test for XError
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

	if testThrowErrorWhenErrorIsNil() != nil {
		t.Fail()
	}

	if testThrowXError() == nil {
		t.Fail()
	}

}

func TestPostHandler(t *testing.T){
	testPostHandlerWithThrowf()
	testPostHandlerWithOtherError()
}

func testPostHandlerWithOtherError() (err error) {
	defer AutoHandle(context.Background(), &err)

	panic("testPostHandlerWithOtherError")
}
func testPostHandlerWithThrowf() (err error) {
	defer AutoHandle(context.Background(), &err)

	Throwf("TestPostHandler")

	return nil
}

func TestXerrorCustomerHandler(t *testing.T) {
	if testCustomerHandler() == nil {
		t.Fail()
	}
}

func testCustomerHandler() (err error) {
	var errorHandlers []ErrorHandler
	errorHandlers = append(errorHandlers, &testHandler{})
	defer Handle(context.Background(), &err, errorHandlers)

	Throwf("testCustomer error happen")

	return nil
}

func testThrowError() (err error) {
	defer AutoHandle(context.Background(), &err)
	e := errors.New("TestThrowError error")

	ThrowError(e)
	return nil
}

func testThrowErrorWhenErrorIsNil() (err error) {
	defer AutoHandle(context.Background(), &err)

	ThrowError(nil)
	return nil
}

func testThrowXError() (err error) {
	defer AutoHandle(context.Background(), &err)

	x := &Xerror{
		errors.New("xerror"),
	}

	ThrowError(x)
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

type testHandler struct{}

func (t testHandler) Match(ctx context.Context, i interface{}) (match bool) {
	return true
}

func (t testHandler) Handle(ctx context.Context, i interface{}) (stop bool) {
	return true
}

func (t testHandler) ConvertToError(ctx context.Context, i interface{}) error {
	return errors.New(fmt.Sprintf("%v", i))
}

type testHandlerNotMatch struct{}

func (t testHandlerNotMatch) Match(ctx context.Context, i interface{}) (match bool) {
	return true
}

func (t testHandlerNotMatch) Handle(ctx context.Context, i interface{}) (stop bool) {
	return false
}

func (t testHandlerNotMatch) ConvertToError(ctx context.Context, i interface{}) error {
	return errors.New(fmt.Sprintf("%v", i))
}

