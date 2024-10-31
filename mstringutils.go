package mstringutils

import (
	"strings"
)

// St is a custom string type.
type St string

// ToUpperCase converts the given string to upper case.
// Example usage:
// s := St("hello")
// result := s.ToUpperCase("world")
// fmt.Println(result) // Output: WORLD
func (s St) ToUpperCase(s1 string) string {
	return strings.ToUpper(s1)
}
