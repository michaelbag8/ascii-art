package main

import (
	//"flag"
	"fmt"
	"os"
	"strings"
	//"strings"
	//"os"
)

var colorMap = map[string]string{
    "red":     "\033[31m",
    "green":   "\033[32m",
    "yellow":  "\033[33m",
    "blue":    "\033[34m",
    "magenta": "\033[35m",
    "cyan":    "\033[36m",
    "white":   "\033[37m",
}

func main() {
	/*
	output := flag.String("output", "", "file to write to")
	color := flag.String("color", "", "colors: red, green, cyan ....")
	align := flag.String("align", "left", "alignment : center, right, left ...")
	reverse := flag.String("reverse", "", "ASCII art to change back to text")
	letters := flag.String("letter", "", "specific letters to color")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Usage: go run . <--flag> <text> <banner>")
		os.Exit(1)
	}

	fmt.Println("output: ", *output)
	fmt.Println("color: ", *color)
	fmt.Println("align: ", *align)
	fmt.Println("reverse: ", *reverse)
	fmt.Println("letters: ", *letters)
	fmt.Println(args)


	var result strings.Builder
	rows := []string{" _  _ ", "| || |", "|_||_|"}
	for _, row := range rows{
		result.WriteString(row)
		result.WriteString("\n")
	}
	fmt.Print(result.String())
	

	content := " _  _ \n| || |\n|_||_|\n"
	err:= os.WriteFile("art.txt", []byte(content), 0644)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Error: writing to file %v\n", err)
		os.Exit(1)
	}

	data, err := os.ReadFile("art.txt")
	if err != nil{
		fmt.Fprintf(os.Stderr, "Error: reading file %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(data))


	const (
    Red     = "\033[31m"
    Green   = "\033[32m"
    Yellow  = "\033[33m"
    Blue    = "\033[34m"
    Magenta = "\033[35m"
    Cyan    = "\033[36m"
    White   = "\033[37m"
    Reset   = "\033[0m"   
)

const (
    Red   = "\033[31m"
    Blue  = "\033[34m"
    Green = "\033[32m"
    Reset = "\033[0m"
)


fmt.Println(Red+ "Hello" + Reset)
fmt.Println(Blue+ "Hello" + Reset)
fmt.Println(Green+ "Hello" + Reset)
fmt.Println("Normal text")


colors := []string{Red, Green, Blue}
for i, ch := range "Hello" {
    color := colors[i%len(colors)]
	fmt.Printf("%s%c%s",color,ch,Reset)
}
fmt.Println()

colorName := "red"
//colorName := "purple"
block, exist := colorMap[colorName]
if !exist{
	fmt.Fprintf(os.Stderr, "Unknown color: %s\n", block)
	os.Exit(1)
}else{
	fmt.Printf("color code for red is \\033[31m\n")
	fmt.Printf("%sHello%s\n",block, "\033[0m")
}
*/

colorCode := "\033[31m"
letters   := "He"
reset     := "\033[0m"

for _, ch := range "Hello World"{
	if letters == "" || strings.ContainsRune(letters, ch){
		fmt.Printf("",colorCode, letters, reset)
	}
}


}
