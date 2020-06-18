package server

import (
	"io/ioutil"
	"net/http"
)

type Server struct {
}

func (s *Server) Start() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8080", nil)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res string
	body, _ := ioutil.ReadAll(r.Body)

	path := r.URL.Path

	switch path {
	case "/HelloService/Hello":
		res = HelloImpl{}.Hello(string(body))
	default:
		res = "not found"
	}

	w.Write([]byte(res))
}
