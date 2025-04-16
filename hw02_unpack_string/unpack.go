package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	runes := []rune(s)
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	var sb strings.Builder
	for i := 0; i < len(runes); i++ {
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			count := int(runes[i+1] - '0')
			if count > 0 {
				sb.WriteString(strings.Repeat(string(runes[i]), count))
			}
			i++

			if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}
		} else {
			sb.WriteRune(runes[i])
		}
	}

	return sb.String(), nil
}
