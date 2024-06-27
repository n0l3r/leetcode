func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var res int
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			res -= romanMap[s[i]]
		} else {
			res += romanMap[s[i]]
		}
	}
	return res
}