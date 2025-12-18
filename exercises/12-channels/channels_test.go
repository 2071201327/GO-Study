package channels

import "testing"

func TestSum(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6} // sum = 21
	if got := Sum(values); got != 21 {
		t.Errorf("Sum() = %d; want 21", got)
	}
}

func TestPingPong(t *testing.T) {
	if got := PingPong(); got != "pong" {
		t.Errorf("PingPong() = %q; want pong", got)
	}
}
