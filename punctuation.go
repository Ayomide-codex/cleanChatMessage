package main

func collapsePunctuation(word string) string {
	result := ""
	for i := 0; i < len(word); i++ {
		if i > 0 && word[i] == word[i-1] && isPunctuation(word[i]) {
			continue
		}
		result += string(word[i])
	}
	return result
}

func isPunctuation(b byte) bool {
	return b == '!' || b == '?' || b == '.' || b == ','
}

func capitalizeFirstLetters(s string) string {
	result := []byte(s)
	capitalizeNext := true

	for i := 0; i < len(result); i++ {
		if capitalizeNext && result[i] >= 'a' && result[i] <= 'z' {
			result[i] -= 32
			capitalizeNext = false
		}
		if isPunctuation(result[i]) {
			capitalizeNext = true
		}
	}
	return string(result)
}
