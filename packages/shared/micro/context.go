package micro

import (
	"context"
	"net/http"
	"time"
)

type Context[TParams any, TInput any] interface {
	context.Context
	RequestID() string
	Header() http.Header
	Params() TParams
	Input() TInput
}

type handlerContext[TParams any, TInput any] struct {
	context   context.Context
	requestID string
	header    http.Header
	params    TParams
	input     TInput
}

var _ Context[any, any] = (*handlerContext[any, any])(nil)

func (c *handlerContext[TParams, TInput]) Value(key any) any {
	return c.context.Value(key)
}

func (c *handlerContext[TParams, TInput]) Deadline() (deadline time.Time, ok bool) {
	return c.context.Deadline()
}

func (c *handlerContext[TParams, TInput]) Done() <-chan struct{} {
	return c.context.Done()
}

func (c *handlerContext[TParams, TInput]) Err() error {
	return c.context.Err()
}

func (c *handlerContext[TParams, TInput]) RequestID() string {
	return c.requestID
}

func (c *handlerContext[TParams, TInput]) Header() http.Header {
	return c.header
}

func (c *handlerContext[TParams, TInput]) Params() TParams {
	return c.params
}

func (c *handlerContext[TParams, TInput]) Input() TInput {
	return c.input
}
