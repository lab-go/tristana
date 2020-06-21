package server

import (
	"context"
)

type ServiceDesc struct {
	ServiceName   string
	ServiceType   interface{}
	MethodMapping []*MethodDesc
}

func (sd *ServiceDesc) GetHandler(methodName string) MethodHandler {
	for _, md := range sd.MethodMapping {
		if md.MethodName == methodName {
			return md.Handler
		}
	}

	return nil
}

type MethodDesc struct {
	MethodName string
	Handler    MethodHandler
}

type MethodHandler func(ctx context.Context, svr interface{}, req []byte) ([]byte, error)
