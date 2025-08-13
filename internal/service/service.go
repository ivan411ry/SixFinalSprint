package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

const allowed = ".- /\t\n\r"

func Convert(input string) (result string, format string, err error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return "", "", errors.New("empty input")
	}

	idx := strings.IndexFunc(s, func(r rune) bool {
		return !strings.ContainsRune(allowed, r)
	})
	if idx == -1 {
		result = morse.ToText(s)
		format = "morse"
	} else {
		result = morse.ToMorse(s)
		format = "text"
	}
	return result, format, nil

}
