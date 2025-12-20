package control

// CheckNumber 判断数字是 "positive", "negative" 还是 "zero"
func CheckNumber(n int) string {
	switch {
	case n > 0:
		return "positive"
	case n < 0:
		return "negative"
	default:
		return "zero"
	}
}

// GetDayName 根据 1-7 返回 "Monday" 等，其他返回 "Unknown"
func GetDayName(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Unknown"
	}
}

// SumUpTo 计算 1 到 n 的和
func SumUpTo(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		result += i
	}

	return result
}

// SumUpToR recursive version of SumUpTo
func SumUpToR(n int) int {
	if n == 0 {
		return 0
	}

	return n + SumUpToR(n-1)
}
