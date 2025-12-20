package variables

// GetName 返回名字 "Gopher"
func GetName() string {
	var name = "Gopher"
	return name
}

// GetAge 返回年龄 10
func GetAge() int {
	var age int = 10
	return age
}

// IsCool 返回是否很酷 (true)
func IsCool() bool {
	cool := true
	return cool
}

// GetZeroValues 返回 int 和 string 的零值
func GetZeroValues() (int, string) {
	var zeroValueInt int
	var zeroValueString string
	return zeroValueInt, zeroValueString
}
