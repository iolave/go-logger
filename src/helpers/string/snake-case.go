package strhelpers

import (
	"strings"
	"unicode"
)

func ToSnakeCase(str string) string {
	var newStr strings.Builder
	var strRune = []rune(str)
	var strLen = len(strRune)

	if strLen == 0 {
		return str
	}

	if unicode.IsLetter(strRune[0]) || unicode.IsNumber(strRune[0]) {
		if unicode.IsUpper(strRune[0]) {
			newStr.WriteRune(unicode.ToLower(strRune[0]))
		} else {
			newStr.WriteRune(strRune[0])
		}
	}

	var prev = strRune[0]
	for _, curr := range strRune[1 : strLen-1] {
		if !unicode.IsLetter(prev) && !unicode.IsNumber(prev) {
			if !unicode.IsLetter(curr) && !unicode.IsNumber(curr) {
				prev = curr
				continue
			}

			newStr.WriteRune(unicode.ToLower(curr))
			prev = curr
			continue
		}

		if !unicode.IsLetter(curr) && !unicode.IsNumber(curr) {
			newStr.WriteRune('_')
			prev = curr
			continue
		}

		if unicode.IsUpper(curr) {
			if !unicode.IsLetter(prev) && !unicode.IsNumber(prev) {
				newStr.WriteRune(unicode.ToLower(curr))
				prev = curr
				continue
			}

			newStr.WriteRune('_')
			newStr.WriteRune(unicode.ToLower(curr))
			prev = curr
			continue
		}

		if unicode.IsLower(curr) {
			if unicode.IsNumber(prev) {
				newStr.WriteRune('_')
			}
			newStr.WriteRune(curr)
			prev = curr
			continue
		}

		newStr.WriteRune(curr)
		prev = curr
	}

	if unicode.IsNumber(strRune[strLen-1]) {
		if !unicode.IsNumber(prev) && !unicode.IsLetter(prev) {
			newStr.WriteRune('_')
			newStr.WriteRune(strRune[strLen-1])
		}

		newStr.WriteRune(strRune[strLen-1])
	}

	if unicode.IsLetter(strRune[strLen-1]) {
		if unicode.IsNumber(prev) {
			newStr.WriteRune('_')
		}

		newStr.WriteRune(unicode.ToLower(strRune[strLen-1]))
	}

	if newStr.String()[newStr.Len()-1:newStr.Len()] == "_" {
		return string(newStr.String())[0 : newStr.Len()-1]
	}

	return string(newStr.String())
}
