package methods

import "testing"

func TestCounter(t *testing.T) {
	c := Counter{Value: 10}

	if val := c.GetValue(); val != 10 {
		t.Errorf("GetValue() = %d; want 10", val)
	}

	c.Increment()
	if c.Value != 11 {
		t.Errorf("After Increment(), Value = %d; want 11", c.Value)
	}
}
