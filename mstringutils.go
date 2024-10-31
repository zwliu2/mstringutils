package mstringutils

import (
	"strings"
)

// St is a custom string type.
type St struct {
	value string
}

// NewSt creates a new St instance with an optional initial value.
func NewSt(value ...string) *St {
	if len(value) > 0 {
		return &St{value: value[0]}
	}
	return &St{value: ""}
}

// ToUpperCase converts the given string to upper case.
func (s *St) ToUpperCase(s1 string) string {
	return strings.ToUpper(s1)
}
