package main

import "unicode"

func collapseRepeatedPunctuation(word string) string {
	var result []rune
	runes := []rune(word)

	for i := 0; i < len(runes); i++ {
		current := runes[i]
		result = append(result, current)

		if isPunctuation(current) {
			for i+1 < len(runes) && runes[i+1] == current {
				i++
			}
		}
	}
	return string(result)
}

func isPunctuation(r rune) bool {
	return r == '!' || r == '?' || r == '.' || r == ','
}

func capitalizeSentences(s string) string {
	runes := []rune(s)
	capitalizeNext := true

	for i, r := range runes {
		if capitalizeNext && unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			capitalizeNext = false
		}

		if isPunctuation(r) {
			capitalizeNext = true
		}
	}
	return string(runes)
}
