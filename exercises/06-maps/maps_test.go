package maps

import "testing"

func TestCreateMap(t *testing.T) {
	m := CreateMap()
	if m["apple"] != 5 || m["banana"] != 10 {
		t.Errorf("CreateMap() returned incorrect map: %v", m)
	}
}

func TestDeleteEntry(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	DeleteEntry(m, "a")
	if _, ok := m["a"]; ok {
		t.Errorf("DeleteEntry failed to remove 'a'")
	}
}

func TestGetScore(t *testing.T) {
	m := map[string]int{"Alice": 100}
	if score, ok := GetScore(m, "Alice"); score != 100 || !ok {
		t.Errorf("GetScore('Alice') = (%d, %t); want (100, true)", score, ok)
	}
	if _, ok := GetScore(m, "Bob"); ok {
		t.Errorf("GetScore('Bob') should return false")
	}
}
