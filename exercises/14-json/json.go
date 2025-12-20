package jsonpkg

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// Encode 将 Response 编码为 JSON
func Encode(r Response) ([]byte, error) {
	// TODO: 实现该方法
	panic("TODO")
}

// Decode 将 JSON 解码为 Response
func Decode(data []byte) (Response, error) {
	// TODO: 实现该方法
	panic("TODO")
}
