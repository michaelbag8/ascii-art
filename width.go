package main

import (
	"golang.org/x/term"
	"os"
)

func getTermWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80
	}
	return width
}
