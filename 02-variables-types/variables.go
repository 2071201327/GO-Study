package variables

func GetName() string {
	var name = "Gopher"
	return name
}

func GetAge() int {
	var age = 10
	return age
}

func IsCool() bool {
	cool := true
	return cool
}

func GetZeroValues() (int, string) {
	var zeroValueInt int
	var zeroValueString string
	return zeroValueInt, zeroValueString
}
