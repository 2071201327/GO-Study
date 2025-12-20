package routing

import "net/http"

type Router struct {
}

// ServeHTTP 实现 http.Handler 接口
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: 实现该方法
	panic("TODO")
}
