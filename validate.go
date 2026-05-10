package main

import (
	"fmt"
)

func Validate(str string) (rune, error) {
	for _, ch := range str {
		if ch < 32 || ch > 126 {
			return ch, fmt.Errorf("not an ascii character %q", ch)
		}
	}
	return 0, nil
}
