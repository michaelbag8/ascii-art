package main

import (
	"fmt"
	"os"
	"strings"
)



func findChar(rows []string, position int, banner map[rune][]string) (rune, int) {
	for ch, charRows := range banner {
		width := len(charRows[0])

		match := true

		for row := 0; row < 8; row++ {
			if position+width > len(rows[row]) {
				match = false
				break
			}

			slice := rows[row][position : position+width]
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


func main(){
	
	data , err := LoadBanner("standard.txt")
	 if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

	content, err := os.ReadFile("reverse_test.txt")
	 if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }

	strings.Split(string(content), "\n")


}