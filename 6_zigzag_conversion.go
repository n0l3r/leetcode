// Source: https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var rows = make([]string, min(numRows, len(s)))
	var curRow int
	var goingDown bool
	for _, c := range s {
		rows[curRow] += string(c)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}
	var ret string
	for _, row := range rows {
		ret += row
	}
	return ret
}