package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	runes := []rune(s)

	if runes[0] >= '0' && runes[0] <= '9' {
		return "", ErrInvalidString
	}

	var sb strings.Builder

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if i+1 < len(runes) && runes[i+1] >= '0' && runes[i+1] <= '9' {
			n, err := strconv.Atoi(string(runes[i+1]))

			if err != nil {
				return "", ErrInvalidString
			}

			if n > 0 {
				sb.WriteString(strings.Repeat(string(r), n))
			}

			i++
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String(), nil
}
