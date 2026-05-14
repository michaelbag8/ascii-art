package main

import (
	"fmt"
	"os"
	"strings"
)

func Generate(input string, banner map[rune][]string, colorCode string, letters string, align string) string {
	if len(input) == 0 {
		return ""
	}
	var output strings.Builder
	segments := SplitInput(input)

	

	for i, segment := range segments {
		if segment == "" {
			if i < len(segments)-1{
				output.WriteString("\n")
			}
			// if i == 0 || segments[i-1] != "" {
			// 	output.WriteString("\n")
			// }
			continue
		}
artWidth := 0
for _, ch := range segment {
    artWidth += len(banner[ch][0])
}

		for _, ch := range segment {
			if _, ok := banner[ch]; !ok {
				fmt.Fprintf(os.Stderr, "Non ASCII character %q\n", ch)
				os.Exit(1)
			}
		}

		rows := renderLines(segment, banner, colorCode, letters, align)
		for _, row := range rows {
			output.WriteString(row)
			output.WriteString("\n")
		}
		if i < len(segments)-1{
				output.WriteString("\n")
			}

	}
	return output.String()
}
