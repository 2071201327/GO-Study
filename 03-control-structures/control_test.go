package control

import "testing"

func TestCheckNumber(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{10, "positive"},
		{-5, "negative"},
		{0, "zero"},
	}

	for _, tt := range tests {
		if got := CheckNumber(tt.input); got != tt.want {
			t.Errorf("CheckNumber(%d) = %q; want %q", tt.input, got, tt.want)
		}
	}
}

func TestGetDayName(t *testing.T) {
	if got := GetDayName(1); got != "Monday" {
		t.Errorf("GetDayName(1) = %q; want Monday", got)
	}
	if got := GetDayName(8); got != "Unknown" {
		t.Errorf("GetDayName(8) = %q; want Unknown", got)
	}
}

func TestSumUpTo(t *testing.T) {
	if got := SumUpTo(5); got != 15 {
		t.Errorf("SumUpTo(5) = %d; want 15", got)
	}
}
