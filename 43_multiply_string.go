
// Source: https://leetcode.com/problems/multiply-strings/
func multiply(num1 string, num2 string) string {
	var res = make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + res[i+j+1]
			res[i+j] += sum / 10
			res[i+j+1] = sum % 10
		}
	}
	var sb strings.Builder
	for i := 0; i < len(res); i++ {
		if !(sb.Len() == 0 && res[i] == 0) {
			sb.WriteString(strconv.Itoa(res[i]))
		}
	}
	if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}