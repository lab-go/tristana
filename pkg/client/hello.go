package client

import (
	"github.com/lab-go/tristana/pkg/share"
	"io/ioutil"
	"net/http"
	"strings"
)

type helloStub struct {
}

func NewHelloService() share.HelloService {
	return &helloStub{}
}

func (h *helloStub) Hello(name string) string {

	res, err := http.Post("http://127.0.0.1:8080/HelloService/Hello", "application/json", strings.NewReader(name))

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}
