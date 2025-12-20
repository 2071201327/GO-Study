package functions

// Calc 返回两个整数的和与差
func Calc(a, b int) (int, int) {
	return a + b, a - b
}

// Swap 交换两个字符串
func Swap(a, b string) (string, string) {
	return b, a
}
