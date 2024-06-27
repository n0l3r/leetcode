func findSubstring(s string, words []string) []int {
	var res []int
	if len(words) == 0 {
		return res
	}
	var wordLen = len(words[0])
	var wordCount = len(words)
	var totalLen = wordLen * wordCount
	var wordMap = make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	for i := 0; i < len(s)-totalLen+1; i++ {
		var seen = make(map[string]int)
		var j int
		for j = 0; j < wordCount; j++ {
			var word = s[i+j*wordLen : i+(j+1)*wordLen]
			if _, ok := wordMap[word]; !ok {
				break
			}
			seen[word]++
			if seen[word] > wordMap[word] {
				break
			}
		}
		if j == wordCount {
			res = append(res, i)
		}
	}
	return res
}