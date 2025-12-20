package functions

// Calc 返回两个整数的和与差
func Calc(a, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

// Swap 交换两个字符串
func Swap(a, b string) (string, string) {
	return b, a
}
