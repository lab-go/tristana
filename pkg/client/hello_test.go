package client

import (
	"fmt"
	"testing"
)

func TestHelloStub_Hello(t *testing.T) {
	fmt.Println(NewHelloService().Hello("world"))
}
