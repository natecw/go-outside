package main

// Boyer-Moore string search
func Search(str, fullText string) int {
	skipTable := preprocess(str)
	skip := 0
	for len(fullText)-skip >= len(str) {
		if same(fullText[skip:], str, len(str)) {
			return skip
		}
		skip = skip + skipTable[fullText[skip+len(str)-1]]
	}
	return -1
}

func preprocess(pattern string) []int {
	t := make([]int, 256)
	for i := 0; i < 256; i++ {
		t[i] = len(pattern)
	}
	for i := 0; i < len(pattern)-1; i++ {
		t[pattern[i]] = len(pattern) - 1 - i
	}
	return t
}

func same(str1, str2 string, len int) bool {
	i := len - 1
	for str1[i] == str2[i] {
		if i == 0 {
			return true
		}
		i -= 1
	}
	return false
}
