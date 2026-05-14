package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	output := flag.String("output", "", "file to write output to")
	color := flag.String("color", "", "colors: red, green, cyan ....")
	letters := flag.String("letters", "", "specific letters to color")
	align := flag.String("align", "left", "alignment: left, right, center, justify")
	_ = flag.String("reverse", "", "ASCII art to change back to text")

	flag.Parse()

	args := flag.Args()

	validAligns := map[string]bool{
		"left": true, "right": true,
		"center": true, "justify": true,
	}
	if *align != "" && !validAligns[*align] {
		fmt.Fprintf(os.Stderr, "unknown alignment: %s\n", *align)
		os.Exit(1)
	}

	colorCode := ""
	if *color != "" {
		code, exist := colorMap[*color]
		if !exist {
			fmt.Fprintf(os.Stderr, "unknown color %s\n", *color)
			os.Exit(1)
		}
		colorCode = code
	}

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "usage go run . [--output=file.txt] <string> [banner]")
		return
	}

	fontFile := "standard"
	input := args[0]

	if len(args) >= 2 {
		fontFile = args[1]
	}

	data, err := LoadBanner(fontFile + ".txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading loadbanner function ")
		return
	}

	content := Generate(input, data, colorCode, *letters, *align)

	if *output != "" {
		err := os.WriteFile(*output, []byte(content), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error : writing file %v\n", err)
			return
		}
		fmt.Printf("Output written to %s\n", *output)
	} else {
		fmt.Print(content)
	}

}
