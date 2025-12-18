package interfaces

import (
	"math"
	"testing"
)

func TestTotalArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 10},
	}

	got := TotalArea(shapes)
	want := 50 + math.Pi*100

	if math.Abs(got-want) > 0.001 {
		t.Errorf("TotalArea() = %f; want %f", got, want)
	}
}
