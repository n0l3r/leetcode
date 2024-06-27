func reverse(x int) int {
	var res int
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res > 1<<31-1 || res < -(1<<31) {
		return 0
	}
	return res
}