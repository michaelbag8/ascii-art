package main

import (
	"strings"
)

func renderLines(segment string, blockMaps map[rune][]string, colorCode string, letters string) []string {
	var output []string

	for row := 0; row < 8; row++ {

		var result strings.Builder
		for _, ch := range segment {
			shouldColor := colorCode != "" && (letters == "" || strings.ContainsRune(letters, ch))
			if shouldColor {
				result.WriteString(colorCode)
				result.WriteString(blockMaps[ch][row])
				result.WriteString("\033[0m")
			} else {
				result.WriteString(blockMaps[ch][row])
			}

		}

		output = append(output, result.String())
	}
	return output
}
