package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var ErrEmptyInput = errors.New("service: пустые входные данные")

const morseAlphabet = ".-/ \n\r"

func DefineOutput(data string) (string, error) {
	trimmed := strings.TrimSpace(data)
	if trimmed == "" {
		return "", ErrEmptyInput
	}

	if isMorseCode(trimmed) {
		return morse.ToText(trimmed), nil
	}

	return morse.ToMorse(trimmed), nil
}

func isMorseCode(s string) bool {
	for _, r := range s {
		if !strings.ContainsRune(morseAlphabet, r) {
			return false
		}
	}
	return true
}
