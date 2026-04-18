package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <text>")
        return
    }

    input := os.Args[1]
    ascii, err := AsciiArt(input)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Print(ascii)
}

func AsciiArt(input string) (string, error) {
    data, err := os.ReadFile("standard.txt")
    if err != nil {
        return "", fmt.Errorf("failed to read font file: %w", err)
    }

    lines := strings.Split(string(data), "\n")

    
    input = strings.ReplaceAll(input, `\n`, "\n")
    input = strings.TrimRight(input, "\n")

    words := strings.Split(input, "\n")
    var result strings.Builder

    for _, word := range words {
        for line := 0; line < 8; line++ {
            for _, char := range word {
                index := int(char- ' ')*9 + 1 + line
                if index >= 0 && index < len(lines) {
                    result.WriteString(lines[index])
                }
            }
            result.WriteByte('\n')
        }
    }

    return result.String(), nil
}
