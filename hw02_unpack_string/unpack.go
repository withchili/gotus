package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	result := ""
	runeMultiplier := -1

	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= '0' && runes[i] <= '9' {
			if runeMultiplier != -1 || i == 0 {
				return "", ErrInvalidString
			}
			runeMultiplier = int(runes[i] - '0')
			continue
		}
		if runeMultiplier == -1 {
			result = string(runes[i]) + result
			continue
		}

		newPart := strings.Repeat(string(runes[i]), runeMultiplier)
		runeMultiplier = -1

		result = newPart + result
	}

	return result, nil
}
