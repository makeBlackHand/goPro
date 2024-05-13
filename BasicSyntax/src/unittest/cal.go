package unittest

func addUpper1(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func sub(n1 int, n2 int) int {
	return n1 - n2
}
