package structs

type Address struct {
	City string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

// NewUser 创建一个新的 User
func NewUser(name string, age int, city string) User {
	// TODO: 实现该方法
	panic("TODO")
}

// UpdateAge 更新 User 的年龄
func UpdateAge(u *User, newAge int) {
	// TODO: 实现该方法
	panic("TODO")
}
