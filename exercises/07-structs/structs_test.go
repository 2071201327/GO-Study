package structs

import "testing"

func TestNewUser(t *testing.T) {
	u := NewUser("Alice", 30, "Wonderland")
	if u.Name != "Alice" || u.Age != 30 || u.Address.City != "Wonderland" {
		t.Errorf("NewUser() = %+v; want {Alice 30 {Wonderland}}", u)
	}
}

func TestUpdateAge(t *testing.T) {
	u := User{Name: "Bob", Age: 20}
	UpdateAge(&u, 21)
	if u.Age != 21 {
		t.Errorf("UpdateAge() failed, age is %d; want 21", u.Age)
	}
}
