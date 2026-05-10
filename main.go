package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3{
		fmt.Fprintf(os.Stderr, "Usage: go run . text banner")
		return
	}

	fontFile := "standard.txt"

	if len(os.Args) == 3{
		fontFile = os.Args[2] + ".txt"
	}

	input := os.Args[1]

	data, err := LoadBanner(fontFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error ")
		return
	}

	content := Generate(input, data)
	fmt.Print(content)

}
