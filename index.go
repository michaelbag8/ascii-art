package main

import (
	"fmt"
	"os"
	"strings"
)

func ParseLoad(filename string) ([][]string, error) {
	blocks := [][]string{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := string(data)

	raw := strings.ReplaceAll(content, "\r\n", "\n")

	rawFile := strings.Split(raw, "\n\n")

	for _, char := range rawFile {
		rows := strings.Split(char, "\n")
		if len(rows) < 8 {
			continue
		}
		blocks = append(blocks, rows[:8])
	}
	return blocks, nil

}

func renderLine(line string, blocks [][]string) {
	var result strings.Builder

	for row := 0; row < 8; row++ {
		result.Reset()

		for _, raw := range line {
			index := int(raw - ' ')
			result.WriteString(blocks[index][row])
		}
		fmt.Println(result.String())
	}

}


func renderAll(input string, blocks [][]string){
	segments := strings.Split(input, "\\n")

	for _, segment := range segments{
		if segment == ""{
			fmt.Println()
		}else{
			renderLine(segment, blocks)
		}
	}
	
}



func main() {
	if len(os.Args) < 2 || len(os.Args) > 3{
		fmt.Println("Usage go run . <text>")
		os.Exit(1)
	}

	fontFile := "standard.txt"

	if len(os.Args) == 3 {
		fontFile = os.Args[2] + ".txt"
	}

	input := os.Args[1]

	blocks, err := ParseLoad(fontFile)
	if err != nil{
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	renderAll(input, blocks)

}

