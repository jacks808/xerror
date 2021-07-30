/**
xerror包, 用来抛出和处理异常.

*/
package xerror

import (
	"context"
	"fmt"
	"sync"
)

var handlers []ErrorHandler
var mutex sync.Mutex

// Xerror
type Xerror struct {
	error
}

// register a new error handler
// each project should call this function when program start.
func Register(h ErrorHandler) {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Printf("\n[xerror] Register a new error handler: %T \n", h)

	handlers = append(handlers, h)
}

// error handler interface
type ErrorHandler interface {
	// Match function
	// this function should judge panic value i is match or not
	// param ctx:       context
	// param recovered: value get from recover()
	// return match:    if return match, xerror will call Handle function
	Match(ctx context.Context, recovered interface{}) (match bool)

	// Handle function
	// this function should handle error, e.g. send alert, write error log, send matrix etc.
	// param ctx:       context
	// param recovered: value get from recover()
	// return stop:     if stop, then handle chain, xerror will call ConvertToError function.
	Handle(ctx context.Context, recovered interface{}) (stop bool)

	// ConvertToError function
	// this function convert your error to error
	// param ctx:       context
	// param recovered: value get from recover()
	// return error:    return an error.
	ConvertToError(ctx context.Context, recovered interface{}) error
}
