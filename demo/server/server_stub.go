package main

import (
	"context"
	"errors"
	"github.com/lab-go/tristana/pkg/server"
)

type HelloImpl struct {
}

func (h *HelloImpl) Hello(name string) string {
	return "hello: " + name
}

var helloDesc = &server.ServiceDesc{
	ServiceName: "HelloService",
	ServiceType: (*HelloImpl)(nil),
	MethodMapping: []*server.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    HelloHandler,
		},
	},
}

func HelloHandler(ctx context.Context, svr interface{}, req []byte) ([]byte, error) {
	s, ok := svr.(*HelloImpl)

	if !ok {
		return nil, errors.New("can't convert to HelloImpl")
	}

	return []byte(s.Hello(string(req))), nil
}
