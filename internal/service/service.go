package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(string string) bool {
	s := strings.TrimSpace(string)
	if s == "" {
		return false
	}

	for _, r := range s {
		if r == '.' || r == '-' || unicode.IsSpace(r) {
			continue
		}
		return false
	}

	if !strings.ContainsAny(s, ".-") {
		return false
	}

	return true
}

func ConvertString(str string) string {
	if isMorse(str) {
		return morse.ToText(str)
	}

	return morse.ToMorse(str)
}
