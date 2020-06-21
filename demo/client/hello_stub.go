package main

import (
	"context"
	"github.com/lab-go/tristana/demo/share"
	"github.com/lab-go/tristana/pkg/client"
)

type helloStub struct {
	c *client.Client
}

func NewHelloService(c *client.Client) share.HelloService {
	return &helloStub{
		c: c,
	}
}

func (h *helloStub) Hello(name string) string {
	var res string
	h.c.Invoke(context.Background(), "/HelloService/Hello", name, &res)

	return res
}

func (h *helloStub) User(c *share.CommonReq) *share.User {
	u := &share.User{}
	h.c.Invoke(context.Background(), "/HelloService/User", c, u)
	return u
}
