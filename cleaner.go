package main

import "strings"

func cleanChatMessage(message string) string {
	words := strings.Fields(message)

	for i, word := range words {
		words[i] = collapseRepeatedPunctuation(word)
	}

	cleaned := strings.Join(words, " ")
	cleaned = capitalizeSentences(cleaned)

	return cleaned
}
