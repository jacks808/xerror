package xerror

import (
	"context"
	"errors"
	"fmt"
	"log"
)

type postHandler struct{}

// match
func (handler postHandler) Match(ctx context.Context, i interface{}) (match bool) {
	return true
}

// handle
func (handler postHandler) Handle(ctx context.Context, i interface{}) (stop bool) {
	log.Printf("Some Error hanppen and catched by PostHandler, ctx: %v, error: %v\n", ctx, i)
	return true
}

// convert to error
func (handler postHandler) ConvertToError(ctx context.Context, i interface{}) error {
	if e, ok := i.(*Xerror); ok {
		return e
	}
	return errors.New(fmt.Sprintf("%v", i))
}

