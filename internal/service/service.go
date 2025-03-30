package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Conv(str string) string {
	if str == "" {
		return morse.ErrNoEncoding{Text: str}.Error()
	}
	isMorse := strings.ContainsFunc(str, func(r rune) bool {
		return r == '-' || r == '.'
	})

	if isMorse {
		return morse.ToText(str)
	} else if strings.ContainsAny(str, "абвгдежзиклмнопрстуфхцчшщъыьэюяАБВГДЕЖЗИКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ0123456789") {
		return morse.ToMorse(str)
	}
	return morse.ErrNoEncoding{Text: str}.Error()
}
