package packages

import "testing"

func TestVisibility(t *testing.T) {
	if got := PublicFunc(); got != "public" {
		t.Errorf("PublicFunc() = %q; want public", got)
	}

	// 在同一个包内（package packages），我们可以访问 privateFunc
	if got := privateFunc(); got != "private" {
		t.Errorf("privateFunc() = %q; want private", got)
	}
}
