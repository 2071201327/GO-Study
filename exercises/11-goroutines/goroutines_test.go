package goroutines

import "testing"

func TestRunConcurrently(t *testing.T) {
	n := 100
	if got := RunConcurrently(n); got != n {
		t.Errorf("RunConcurrently(%d) = %d; want %d", n, got, n)
	}
}
