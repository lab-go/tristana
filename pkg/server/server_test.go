package server

import "testing"

func TestServeHTTP(t *testing.T) {
	server := Server{}
	server.Start()
}
