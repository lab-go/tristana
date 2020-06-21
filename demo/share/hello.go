package share

type HelloService interface {
	Hello(name string) string
	User(*CommonReq) *User
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CommonReq struct {
	Common string `json:"common"`
}
