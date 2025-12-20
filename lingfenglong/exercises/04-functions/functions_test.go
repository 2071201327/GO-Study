package functions

import "testing"

func TestCalc(t *testing.T) {
	sum, diff := Calc(10, 5)
	if sum != 15 || diff != 5 {
		t.Errorf("Calc(10, 5) = (%d, %d); want (15, 5)", sum, diff)
	}
}

func TestSwap(t *testing.T) {
	a, b := Swap("hello", "world")
	if a != "world" || b != "hello" {
		t.Errorf("Swap(\"hello\", \"world\") = (%q, %q); want (\"world\", \"hello\")", a, b)
	}
}
