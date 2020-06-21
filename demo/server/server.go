package main

import "github.com/lab-go/tristana/pkg/server"

func main() {
	server := server.NewServer()
	server.RegisterService(helloDesc, &HelloImpl{})
	server.Start()
}
