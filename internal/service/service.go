package service

import (
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Conv(str string) string {
	var result string
	if str == "" {
		fmt.Println("empty text, nothing to convert")
		return ""
	}
	for _, ch := range str {
		if ch == 45 || ch == 46 {
			result = morse.ToText(str)
			return result
		}
		result = morse.ToMorse(str)
		return result
	}
	return result
}
