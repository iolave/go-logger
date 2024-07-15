package strutils

import (
	"strings"
	"unicode"
)

func ToSnakeCase(str string) string {
	var newStr strings.Builder
	var strRune = []rune(str)
	var strLen = len(strRune)

	// If the string length is 0 we know we don't have to snake case
	// anything and therefore we'll simply return the empty string.
	if strLen == 0 {
		return str
	}

	// If the first rune is either a letter or a number
	// we will add it to our result string.
	// If a letter happens to be upper case, we add the
	// lower case version of it.
	if !(!unicode.IsLetter(strRune[0]) && !unicode.IsNumber(strRune[0])) {
		newStr.WriteRune(unicode.ToLower(strRune[0]))
	}

	// If the string length is 1 simply check if the
	// rune is a symbol. If it is, return an empty string.
	// Otherwise return the lower case version of the rune.
	if strLen == 1 {
		if !unicode.IsLetter(strRune[0]) && !unicode.IsNumber(strRune[0]) {
			return ""
		}

		return string(unicode.ToLower(strRune[0]))
	}

	// Iterates over the rest of the runes but
	// the last one.
	var prev = strRune[0]
	for _, curr := range strRune[1 : strLen-1] {
		isPrevLetter := unicode.IsLetter(prev)
		isPrevNumber := unicode.IsNumber(prev)
		isPrevSymbol := !isPrevLetter && !isPrevNumber
		isCurrNumber := unicode.IsNumber(curr)
		isCurrSymbol := !unicode.IsLetter(curr) && !isCurrNumber

		// When the current rune is a symbol it's value
		// is replaced with an underscore, only if the
		// previous rune is a not a symbol cuz it means
		// that we already have an underscore rune.
		if isCurrSymbol {
			if isPrevSymbol {
				prev = curr
				continue
			}

			newStr.WriteRune('_')
			prev = curr
			continue
		}

		// When a number rune is found we will
		// add it to the result stringf, but before
		// that we add an underscore rune to the
		// result stringf only if the prev rune is
		// a letter.
		if isCurrNumber {
			if isPrevLetter {
				newStr.WriteRune('_')
			}

			newStr.WriteRune(curr)
			prev = curr
			continue
		}

		// If the current rune is upper case we will lower
		// case it, append an underscore to the result string
		// only if the previous rune is not a symbol, append
		// the lower case version of it to the result string
		// and add another underscore rune after.
		if unicode.IsUpper(curr) {
			if isPrevNumber {
				newStr.WriteRune('_')
			}

			newStr.WriteRune(unicode.ToLower(curr))
			newStr.WriteRune('_')
			prev = '_'
			continue
		}

		// At this point, the current rune is in fact a lower
		// case letter.
		// We add an underscore rune to the result string but
		// only if the previous rune is a number and finally
		// write the current rune to the result
		if isPrevNumber {
			newStr.WriteRune('_')
		}
		newStr.WriteRune(curr)
		prev = curr
	}

	// Finally, we handle the last rune
	isPrevLetter := unicode.IsLetter(prev)
	isPrevNumber := unicode.IsNumber(prev)
	isPrevSymbol := !isPrevLetter && !isPrevNumber
	isLastNumber := unicode.IsNumber(strRune[strLen-1])
	isLastSymbol := !unicode.IsLetter(strRune[strLen-1]) && !isLastNumber

	// If the last rune is a symbol we skip this rune
	// and we will check if the prev rune was a symbol.
	// If this is true, we remove the last added rune
	// of the result string.
	if isLastSymbol {
		if isPrevSymbol {
			return string(newStr.String()[0 : newStr.Len()-1])
		}
		return string(newStr.String())
	}

	// If the last rune happens to be a number,
	// we add an underscore rune to the result string
	// only if the previous rune is a letter.
	// Finally, add the last rune to the result string.
	if isLastNumber {
		if isPrevLetter {
			newStr.WriteRune('_')
		}

		newStr.WriteRune(strRune[strLen-1])
		return string(newStr.String())
	}

	// At this point, the last rune is in fact a letter.
	// If the previous rune is a number add an underscore
	// rune to the result string and finally, add the lower
	// case version of the letter to the result string.
	if isPrevNumber {
		newStr.WriteRune('_')
	}

	// If the last rune is upper case add an underscore
	// rune to the result string and finally, add the lower
	// case version of the letter to the result string.
	if unicode.IsUpper(strRune[strLen-1]) {
		newStr.WriteRune('_')
	}

	newStr.WriteRune(unicode.ToLower(strRune[strLen-1]))

	return string(newStr.String())
}
