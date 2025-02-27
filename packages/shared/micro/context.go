package micro

import "net/http"

type Context[TParams any, TInput any] interface {
	Header() http.Header
	Params() TParams
	Input() TInput
}

type handlerContext[TParams any, TInput any] struct {
	header http.Header
	params TParams
	input  TInput
}

var _ Context[any, any] = (*handlerContext[any, any])(nil)

func (c *handlerContext[TParams, TInput]) Header() http.Header {
	return c.header
}

func (c *handlerContext[TParams, TInput]) Params() TParams {
	return c.params
}

func (c *handlerContext[TParams, TInput]) Input() TInput {
	return c.input
}
