package main

import "strings"

func main() {
	println(findFirstStringInBracket("sda(qqwew)asd"))
}

func findFirstStringInBracket(input string) string {
	// Validate len of input
	if len(input) == 0 {
		return ""
	}

	// Get index of open bracket, and validate it
	indexOpenBracketFound := strings.Index(input, "(")
	if indexOpenBracketFound == -1 {
		return ""
	}

	// Get index of closing bracket, and validate it
	indexClosingBracketFound := strings.Index(input, ")")
	if indexClosingBracketFound == -1 {
		return ""
	}

	inputRunes := []rune(input)
	return string(inputRunes[indexOpenBracketFound+1 : indexClosingBracketFound])
}
