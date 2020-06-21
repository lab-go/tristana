package main

import (
	"fmt"
	"github.com/lab-go/tristana/demo/share"
	"github.com/lab-go/tristana/pkg/client"
	"net/http"
)

func main() {

	c := client.NewClient(&http.Client{}, "http://127.0.0.1:8080")
	fmt.Println(NewHelloService(c).Hello("world"))

	fmt.Println(NewHelloService(c).User(&share.CommonReq{Common: "Tom"}))
}
