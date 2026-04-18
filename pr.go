package main

import "fmt"

func main() {
	font := map[rune][]string{
		'H': {
			" _    _ ",
			"| |  | |",
			"| |__| |",
			"|  __  |",
			"| |  | |",
			"| |  | |",
			"|_|  |_|",
			"        ",
		},
		'E': {
			" ______ ",
			"|  ____|",
			"| |__   ",
			"|  __|  ",
			"| |____ ",
			"|______|",
			"        ",
			"        ",
		},
		'L': {
			" _      ",
			"| |     ",
			"| |     ",
			"| |     ",
			"| |____ ",
			"|______|",
			"        ",
			"        ",
		},
		'O': {
			"  ____  ",
			" / __ \\ ",
			"| |  | |",
			"| |  | |",
			"| |__| |",
			" \\____/ ",
			"        ",
			"        ",
		},
	}

	input := "HELLO"

	height := 8

	for i := 0; i < height; i++ {
		for _, ch := range input {
			if letter, ok := font[ch]; ok {
				fmt.Print(letter[i])
			} else {
				fmt.Print("        ")
			}
		}
		fmt.Println()
	}
}
