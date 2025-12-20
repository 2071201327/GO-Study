package slices

import (
	"reflect"
	"testing"
)

func TestCreateSlice(t *testing.T) {
	want := []int{1, 2, 3}
	if got := CreateSlice(); !reflect.DeepEqual(got, want) {
		t.Errorf("CreateSlice() = %v; want %v", got, want)
	}
}

func TestAddElement(t *testing.T) {
	s := []int{1, 2}
	want := []int{1, 2, 3}
	if got := AddElement(s, 3); !reflect.DeepEqual(got, want) {
		t.Errorf("AddElement() = %v; want %v", got, want)
	}
}

func TestGetSubSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	want := []int{2, 3} // index 1 to 3(exc)
	if got := GetSubSlice(s, 1, 3); !reflect.DeepEqual(got, want) {
		t.Errorf("GetSubSlice() = %v; want %v", got, want)
	}
}
