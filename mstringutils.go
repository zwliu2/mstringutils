package mstringutils

import (
	"fmt"
	"strings"
)

// St is a custom string type.
type St struct {
	value string
}

// NewSt creates a new St instance with an optional initial value.
func NewSt(value string) (*St, error) {
	if len(value) > 1024 { // 假设最大长度为1024
		return nil, fmt.Errorf("input string is too long")
	}
	return &St{value: value}, nil
}

// ToUpperCase converts the given string to upper case.
func (s *St) ToUpperCase(s1 string) string {
	if s1 == "" {
		return ""
	}
	return strings.ToUpper(s1)
}
