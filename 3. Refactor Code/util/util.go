package util

import "strings"

func findFirstStringInBracket(str string) string {
	if len(str) == 0 {
		return ""
	}
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	wordsAfterFirstBracket := str[indexFirstBracketFound:]
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	return wordsAfterFirstBracket[1:indexClosingBracketFound]
}
