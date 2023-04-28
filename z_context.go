package gin

import (
	"context"
	"time"
)

func (c *Context) Context() context.Context {
	return c.Request.Context()
}

func (c *Context) WithValue(key, val any) context.Context {
	ctx := context.WithValue(c.Request.Context(), key, val)
	c.Request = c.Request.WithContext(ctx)
	return ctx
}

func (c *Context) WithTimeout(timeout time.Duration) context.Context {
	ctx, _ := context.WithTimeout(c.Request.Context(), timeout)
	c.Request = c.Request.WithContext(ctx)
	return c.Request.Context()
}
