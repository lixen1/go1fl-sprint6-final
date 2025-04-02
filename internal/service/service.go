package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DataConvert(data string) string {
	isMorse := strings.Trim(data, ".- ") == ""

	if isMorse {
		return morse.ToText(data)
	}
	return morse.ToMorse(data)
}
