package server

import (
	"context"
	"encoding/json"
	"github.com/lab-go/tristana/pkg/server/codec"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type server struct {
	sMux       sync.RWMutex
	svcMapping map[string]*ServiceDesc
}

func NewServer() *server {
	return &server{
		sMux:       sync.RWMutex{},
		svcMapping: map[string]*ServiceDesc{},
	}
}

func (s *server) RegisterService(svc *ServiceDesc, svr interface{}) {
	s.sMux.Lock()
	svc.ServiceType = svr
	s.svcMapping[svc.ServiceName] = svc
	s.sMux.Unlock()
}

func (s *server) getService(svcName string) *ServiceDesc {
	s.sMux.RLock()
	defer s.sMux.RUnlock()
	return s.svcMapping[svcName]
}

func (s *server) Start() {
	http.ListenAndServe("0.0.0.0:8080", s)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	path := r.URL.Path

	split := strings.Split(path, "/")

	if len(split) != 3 {
		w.Write([]byte("PATH NOT SUPPORT"))
	}

	if svc := s.getService(split[1]); svc == nil {
		w.Write([]byte("NO SERVICE FOUND"))
	} else {
		if handler := svc.GetHandler(split[2]); handler == nil {
			w.Write([]byte("NO METHOD FOUND"))
		} else {
			res, err := handler(context.Background(), svc.ServiceType, body, codec.JsonDecode)

			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				marshal, err := json.Marshal(res)

				if err != nil {
					w.Write([]byte(err.Error()))
				}

				w.Write(marshal)
			}
		}
	}
}
