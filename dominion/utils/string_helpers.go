package utils

import "strconv"

func IsLetter(input string) bool {
	if len(input) == 0 {
		return false
	}

	return IsLetterByte(input[0])
}

func IsLetterByte(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func IsWhitespace(input string) bool {
	if len(input) == 0 {
		return true
	}

	return IsWhiteSpaceByte(input[0])
}

func IsWhiteSpaceByte(input byte) bool {
	return input == 32 || input == 9
}

func IsSymbol(input string) bool {
	if len(input) == 0 {
		return false
	}

	return !IsLetter(input) && !IsWhitespace(input) && !IsNumber(input)
}

func IsNumber(input string) bool {
	_, err := strconv.Atoi(input)

	return err == nil
}

func GetReadUntilPredicate(predicate func(string) bool) func(string) string {
	return func(input string) string {
		for index := range len(input) {
			current := (input)[index]
			if !predicate(string(current)) {
				return (input)[:index]
			}
		}

		return input
	}
}
