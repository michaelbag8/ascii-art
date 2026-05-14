// align.go
package main

import "strings"

func alignText(art string, align string, termWidth, artWidth int) string {
    switch align {
    case "right":
        padding := strings.Repeat(" ", termWidth-artWidth)
        return padding + art
    case "center":
        padding := strings.Repeat(" ", (termWidth-artWidth)/2)
        return padding + art
    case "justify":
        return art
    default:
        return art
    }
}