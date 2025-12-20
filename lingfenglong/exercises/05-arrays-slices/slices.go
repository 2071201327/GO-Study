package slices

// CreateSlice 创建一个包含 1, 2, 3 的切片
func CreateSlice() []int {
	return []int{1, 2, 3}
}

// AddElement 向切片添加一个元素
func AddElement(s []int, n int) []int {
	return append(s, n)
}

// GetSubSlice 获取切片的子切片 [start:end)
func GetSubSlice(s []int, start, end int) []int {
	return s[start:end]
}
