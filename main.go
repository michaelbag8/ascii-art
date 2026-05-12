package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	output := flag.String("output", "", "file to write output to")
	flag.Parse()

	args := flag.Args()

	if len(args) < 1{
		fmt.Fprintf(os.Stderr, "usage go run . [--output=file.txt] <string> [banner]")
		return
	}

	fontFile := "standard"
	input := args[0]

	if len(args) >=2{
		fontFile = args[1]
	}

	data, err := LoadBanner(fontFile + ".txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading loadbanner function ")
		return
	}

	content := Generate(input, data)

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
