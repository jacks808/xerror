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
func handle(ctx context.Context, err *error, handlers []ErrorHandler, recovered interface{}) {
	// got recovered value
	if recovered != nil {
		// ask each handler
		for _, handler := range handlers {
			// if doesn't match, continue
			if !handler.Match(ctx, recovered) {
				continue
			}
			// if match, call handler.ConvertToError to convert recovered to err.
			if stop := handler.Handle(ctx, recovered); stop {
				*err = handler.ConvertToError(ctx, recovered)
				return
			}
		}
		// panic if none of one handlers could handle your error
		panic(recovered)
	}
}
