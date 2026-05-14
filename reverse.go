package main

import(
	"os"
	"fmt"
	"strings"
)

func findChar(rows []string, positionition int, banner map[rune][]string) (rune, int) {
	for ch, charRows := range banner {
		width := len(charRows[0])

		match := true

		for row := 0; row < 8; row++ {
			if positionition+width > len(rows[row]) {
				match = false
				break
			}

			slice := rows[row][positionition : positionition+width]
			if slice != charRows[row] {
				match = false
				break
			}

		}
		if match {
			return ch, width
		}
	}
	return 0, 0
}


func Reverse(filename string, charMap map[rune][]string) (string, error) {
    
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("error reading file: %w", err)
    }

    lines := strings.Split(string(data), "\n")
    var result strings.Builder

   
    for i := 0; i+8 <= len(lines); i += 9 {
        block := lines[i : i+8]

       
        position := 0
        for position < len(block[0]) {
            ch, width := findChar(block, position, charMap)
            if ch == 0 {
                break
            }
            result.WriteString(string(ch))
            position += width
        }


        if i+9 < len(lines) {
            result.WriteString("\\n")
        }
    }

    return result.String(), nil
}