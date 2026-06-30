package main

import "strings"

func cleanChatMessage(message string) string {
	words := strings.Fields(message)

	for i, word := range words {
		words[i] = collapsePunctuation(word)
	}

	cleaned := strings.Join(words, " ")
	cleaned = capitalizeFirstLetters(cleaned)
	return cleaned
}
