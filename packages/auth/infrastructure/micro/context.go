package micro

import "net/http"

type Context[T any] interface {
	Header() http.Header
	Input() T
}

type handlerContext[T any] struct {
	header http.Header
	input  T
}

var _ Context[any] = (*handlerContext[any])(nil)

func (c *handlerContext[T]) Header() http.Header {
	return c.header
}

func (c *handlerContext[T]) Input() T {
	return c.input
}
