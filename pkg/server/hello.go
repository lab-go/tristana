package server

type HelloImpl struct {
}

func (h HelloImpl) Hello(name string) string {
	return "hello: " + name
}
