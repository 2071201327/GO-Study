package control

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

func GetDayName(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 7:
		return "Sunday"
	default:
		return "Unknown"
	}
}

func SumUpTo(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func SumUpToR(n int) int {
	if n == 0 {
		return 0
	}
	return n + SumUpToR(n-1)
}
