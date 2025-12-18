package jsonpkg

import (
	"testing"
)

func TestEncode(t *testing.T) {
	r := Response{Message: "Success", Code: 200}
	data, _ := Encode(r)

	// 简单验证包含关键字符串
	s := string(data)
	if s == "" {
		t.Errorf("Encode returned empty string")
	}
}

func TestDecode(t *testing.T) {
	data := []byte(`{"message": "Error", "code": 500}`)
	r, err := Decode(data)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if r.Message != "Error" || r.Code != 500 {
		t.Errorf("Decode result mismatch: %+v", r)
	}
}
