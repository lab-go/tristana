package main

import (
	"context"
	"errors"
	"github.com/lab-go/tristana/demo/share"
	"github.com/lab-go/tristana/pkg/server"
)

type HelloImpl struct {
}

func (h *HelloImpl) User(common *share.CommonReq) *share.User {
	return &share.User{
		Name: common.Common,
		Age:  10,
	}
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
		{
			MethodName: "User",
			Handler:    UserHandler,
		},
	},
}

func HelloHandler(ctx context.Context, svr interface{}, req []byte, decode server.Decode) (interface{}, error) {
	s, ok := svr.(*HelloImpl)

	if !ok {
		return nil, errors.New("can't convert to HelloImpl")
	}

	var r string

	if err := decode(req, &r); err != nil {
		return nil, err
	}

	return s.Hello(r), nil
}

func UserHandler(ctx context.Context, svr interface{}, req []byte, decode server.Decode) (interface{}, error) {
	s, ok := svr.(*HelloImpl)

	if !ok {
		return nil, errors.New("can't convert to HelloImpl")
	}

	r := &share.CommonReq{}

	if err := decode(req, r); err != nil {
		return nil, err
	}

	return s.User(r), nil
}
