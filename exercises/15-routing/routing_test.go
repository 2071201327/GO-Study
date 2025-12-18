package routing

import (
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	router := &Router{}

	tests := []struct {
		path string
		want string
		code int
	}{
		{"/", "Home", 200},
		{"/about", "About", 200},
		{"/unknown", "Not Found", 404},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.path, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != tt.code {
			t.Errorf("Path %q code = %d; want %d", tt.path, w.Code, tt.code)
		}
		if w.Body.String() != tt.want {
			t.Errorf("Path %q body = %q; want %q", tt.path, w.Body.String(), tt.want)
		}
	}
}
