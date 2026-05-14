package main

import (
    "os"
    "strings"
    "golang.org/x/term"
)

func getTermWidth() int {
    width, _, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        return 80
    }
    return width
}

func renderLines(segment string, blockMaps map[rune][]string, colorCode string, letters string, align string) []string {
    var output []string
    termWidth := getTermWidth()
    chars     := []rune(segment)

    // calculate artWidth from row 0
    artWidth := 0
    for _, ch := range chars {
        artWidth += len(blockMaps[ch][0])
    }

    if align == "justify" && len(chars) > 1 {
        remainingSpace := termWidth - artWidth
        gaps           := len(chars) - 1
        spacePerGap    := remainingSpace / gaps
        extraSpaces    := remainingSpace % gaps

        for row := 0; row < 8; row++ {
            var rawRow strings.Builder
            for i, ch := range chars {
                
                shouldColor := colorCode != "" && (letters == "" || strings.ContainsRune(letters, ch))
                if shouldColor {
                    rawRow.WriteString(colorCode)
                    rawRow.WriteString(blockMaps[ch][row])
                    rawRow.WriteString("\033[0m")
                } else {
                    rawRow.WriteString(blockMaps[ch][row])
                }
                
                if i < len(chars)-1 {
                    gapSize := spacePerGap
                    if i < extraSpaces {
                        gapSize++
                    }
                    rawRow.WriteString(strings.Repeat(" ", gapSize))
                }
            }
            output = append(output, rawRow.String())
        }
        return output
    }

    // left, right, center
    for row := 0; row < 8; row++ {
        var rawRow strings.Builder
        for _, ch := range chars {
            shouldColor := colorCode != "" && (letters == "" || strings.ContainsRune(letters, ch))
            if shouldColor {
                rawRow.WriteString(colorCode)
                rawRow.WriteString(blockMaps[ch][row])
                rawRow.WriteString("\033[0m")
            } else {
                rawRow.WriteString(blockMaps[ch][row])
            }
        }
        output = append(output, alignText(rawRow.String(), align, termWidth, artWidth))
    }
    return output
}