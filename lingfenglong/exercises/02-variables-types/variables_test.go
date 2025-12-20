package variables

import "testing"

func TestVariables(t *testing.T) {
	if got := GetName(); got != "Gopher" {
		t.Errorf("GetName() = %q; want %q", got, "Gopher")
	}
	if got := GetAge(); got != 10 {
		t.Errorf("GetAge() = %d; want %d", got, 10)
	}
	if got := IsCool(); !got {
		t.Errorf("IsCool() = %v; want true", got)
	}
}

func TestZeroValues(t *testing.T) {
	i, s := GetZeroValues()
	if i != 0 {
		t.Errorf("GetZeroValues() int = %d; want 0", i)
	}
	if s != "" {
		t.Errorf("GetZeroValues() string = %q; want \"\"", s)
	}
}
