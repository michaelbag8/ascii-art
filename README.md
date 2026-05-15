# ASCII Art Generator

## Overview

The ASCII Art Generator is a Go application that converts normal text into ASCII art using banner templates. The program supports multiple CLI features such as color formatting, alignment, file output, reverse ASCII decoding, and a web interface mode.

This project demonstrates:

* File handling in Go
* String manipulation
* Error handling
* Maps and slices
* Unit testing
* Command-line flag parsing
* ANSI color formatting
* Terminal width handling
* HTTP server development in Go

---

# Features

* Convert plain text into ASCII art
* Support for banner templates
* Multi-line input support
* Output redirection into files
* ASCII art color formatting
* Text alignment support
* Reverse ASCII art decoding
* Web server interface
* Error handling for invalid input and missing files
* Unit tests for core functions

---

# Project Structure

```text
ascii-art/
├── main.go
├── banner.go
├── validate.go
├── split.go
├── render.go
├── reverse.go
├── align.go
├── color.go
├── output.go
├── server.go
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
├── validate_test.go
├── split_test.go
├── render_test.go
└── README.md
```

---

# Requirements

* Go 1.20 or later

Check your Go version:

```bash
go version
```

---

# Installation

Clone the repository:

```bash
git clone https://github.com/michaelbag8/ascii-art.git
cd ascii-art
```

---

# Running the Project

## Normal Mode

```bash
go run . "Hello"
```

Example output:

```text
 _   _          _   _
| | | |   ___  | | | |
| |_| |  / _ \ | | | |
|  _  | |  __/ | | | |
|_| |_|  \___| |_| |_|
```

---

# Using Different Banner Files

You can use different banner styles:

```bash
go run . "Hello" shadow
```

or

```bash
go run . "Hello" thinkertoy
```

Available banner files:

* standard.txt
* shadow.txt
* thinkertoy.txt

---

# Command-Line Flags

| Flag                  | Description                                          |
| --------------------- | ---------------------------------------------------- |
| `--color=<color>`     | Changes ASCII art color                              |
| `--output=<file>`     | Saves output into a file                             |
| `--align=<type>`      | Aligns output (`left`, `right`, `center`, `justify`) |
| `--letters=<letters>` | Applies color to specific letters                    |
| `--reverse=<file>`    | Converts ASCII art back to text                      |

---

# Usage Examples

## Color Mode

```bash
go run . --color=red "Hello"
```

---

## Align Mode

```bash
go run . --align=center "Hello"
```

Available alignments:

* left
* right
* center
* justify

---

## Output Mode

Save ASCII art into a file:

```bash
go run . --output=test.txt "Hello"
```

Read the saved output:

```bash
cat test.txt
```

---

## Reverse Mode

Generate ASCII art into a file:

```bash
go run . --output=test.txt "Hello\nWorld"
```

Convert ASCII art back into text:

```bash
go run . --reverse=test.txt
```

---

# Web Server Mode (fs)

Start the web server:

```bash
go run . fs
```

Then open your browser and test the form with:

* Text input
* Color selection
* Alignment options
* Banner selection

The server uses Go's `net/http` package to render ASCII art in the browser.

---

# Testing

Run all tests:

```bash
go test
```

Run tests with verbose output:

```bash
go test -v
```

---

# Core Functions

## Validate

Checks whether all characters in the input string are printable ASCII characters.

```go
func Validate(str string) (rune, error)
```

---

## SplitInput

Splits input text using escaped newline characters.

```go
func SplitInput(str string) []string
```

---

## renderLines

Renders ASCII art lines using the selected banner map.

```go
func renderLines(segment string, blockMaps map[rune][]string) []string
```

---

## ReverseASCII

Converts ASCII art back into normal text.

```go
func ReverseASCII(file string) (string, error)
```

---

## AlignText

Handles text alignment inside the terminal.

```go
func AlignText(lines []string, align string) []string
```

---

## startServer

Starts the HTTP web server.

```go
func startServer()
```

---

# Feature Breakdown

## Feature 1 — `--output`

Concepts used:

* `strings.Builder`
* `os.WriteFile`
* `flag.String`
* `flag.Parse`
* `flag.Args()` vs `os.Args`
* Pointers and dereferencing

Example:

```bash
go run . --output=result.txt "Hello"
```

---

## Feature 2 — `--color`

Concepts used:

* ANSI escape codes
* Color sandwich formatting:

```go
\033[31m...\033[0m
```

* `colorMap`
* `strings.ContainsRune`
* `shouldColor` logic

Example:

```bash
go run . --color=blue "Hello"
```

---

## Feature 3 — `--align`

Concepts used:

* Terminal width detection using:

```go
golang.org/x/term
```

* Left/right/center padding
* `strings.Repeat`
* Justify gap calculation
* Modulo distribution logic

Example:

```bash
go run . --align=justify "Hello World"
```

---

## Feature 4 — `--reverse`

Concepts used:

* `findChar` slice matching algorithm
* Block stepping:

```go
i += 9
```

* Multi-block recovery using newline separators

Example:

```bash
go run . --reverse=test.txt
```

---

## Feature 5 — `fs`

Concepts used:

* `net/http`
* `HandleFunc`
* `ListenAndServe`
* `http.ResponseWriter`
* `r.FormValue`
* `Content-Type` headers
* CSS colors vs ANSI colors
* `startServer()`
* `fs` detection in `main()`

Example:

```bash
go run . fs
```

---

# Error Handling

The project handles:

* Invalid ASCII characters
* Missing banner files
* Incorrect input
* Invalid alignment types
* Invalid color values
* File read errors

Example:

```text
Error reading file: open notfound.txt: no such file or directory
```

---

# Learning Objectives

This project helps you practice:

* Go syntax and structure
* Loops and conditionals
* Working with maps and slices
* File operations
* CLI flag parsing
* ANSI terminal formatting
* HTTP server development
* Writing reusable functions
* Unit testing in Go

---

# Future Improvements

Possible enhancements:

* Add more banner styles
* Add downloadable web output
* Improve reverse detection accuracy
* Add Docker support
* Add live preview in browser
* Improve CLI argument validation

---

# Author

Michael BAG

---
