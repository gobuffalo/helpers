package contexth

import (
	"context"

	"github.com/gobuffalo/helpers/hctx"
)

const (
	BackgroundKey = "Background"

	TODOKey = "TODO"

	WithCancelKey = "WithCancel"

	WithDeadlineKey = "WithDeadline"

	WithTimeoutKey = "WithTimeout"

	WithValueKey = "WithValue"
)

func New() hctx.Map {
	return hctx.Map{

		BackgroundKey: Background,

		TODOKey: TODO,

		WithCancelKey: WithCancel,

		WithDeadlineKey: WithDeadline,

		WithTimeoutKey: WithTimeout,

		WithValueKey: WithValue,
	}
}

// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
var Background = context.Background

// TODO returns a non-nil, empty Context. Code should use context.TODO when
// it&#39;s unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter). TODO is recognized by static analysis tools that determine
// whether Contexts are propagated correctly in a program.
var TODO = context.TODO

// WithCancel returns a copy of parent with a new Done channel. The returned
// context&#39;s Done channel is closed when the returned cancel function is called
// or when the parent context&#39;s Done channel is closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
var WithCancel = context.WithCancel

// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent&#39;s deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// context&#39;s Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context&#39;s Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
var WithDeadline = context.WithDeadline

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete:
//
// 	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
// 		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
// 		defer cancel()  // releases resources if slowOperation completes before timeout elapses
// 		return slowOperation(ctx)
// 	}
var WithTimeout = context.WithTimeout

// WithValue returns a copy of parent in which the value associated with key is
// val.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The provided key must be comparable and should not be of type
// string or any other built-in type to avoid collisions between
// packages using context. Users of WithValue should define their own
// types for keys. To avoid allocating when assigning to an
// interface{}, context keys often have concrete type
// struct{}. Alternatively, exported context key variables&#39; static
// type should be a pointer or interface.
var WithValue = context.WithValue
