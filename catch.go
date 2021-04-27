package xerror

import (
	"context"
)

var post postHandler

// handle error with exists error handlers
func AutoHandle(ctx context.Context, err *error) {
	handle(ctx, err, append(handlers, post), recover())
}

// handle error with handlers param
func Handle(ctx context.Context, err *error, handlers []ErrorHandler) {
	handle(ctx, err, handlers, recover())
}

// handler panic via responsibility chain
func handle(ctx context.Context, err *error, handlers []ErrorHandler, r interface{}) {
	if r != nil {
		for _, handler := range handlers {
			if !handler.Match(ctx, r) {
				continue
			}
			if stop := handler.Handle(ctx, r); stop {
				*err = handler.ConvertToError(ctx, r)
				return
			}
		}
		panic(r)
	}
}
