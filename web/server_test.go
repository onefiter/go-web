package web

import "testing"

func TestServer(t *testing.T) {
	s := NewHTTPServer()
	s.Get("/", func(context *Context) {
		context.Resp.Write([]byte("hello world"))
	})

	s.Get("/user", func(context *Context) {
		context.Resp.Write([]byte("hello user"))
	})
	s.Start(":8081")
}
