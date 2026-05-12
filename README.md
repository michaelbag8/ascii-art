# ASCII Art Generator

## Overview

The ASCII Art Generator is a Go application that converts normal text into ASCII art using banner templates. The program reads a banner file containing predefined character patterns and renders the input text into stylized ASCII output.

This project demonstrates:

* File handling in Go
* String manipulation
* Error handling
* Maps and slices
* Unit testing
* Working with command-line arguments

---

# Features

* Convert plain text into ASCII art
* Support for banner templates
* ASCII character validation
* Multi-line input support
* Error handling for invalid files and characters
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

Run the application with:

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

Make sure the corresponding banner file exists:

* standard.txt
* shadow.txt
* thinkertoy.txt

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

# Error Handling

The project handles:

* Invalid ASCII characters
* Missing banner files
* Incorrect input
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
* Writing reusable functions
* Unit testing in Go

---

| Flag                  | Description                                          |
| --------------------- | ---------------------------------------------------- |
| `--color=<color>`     | Changes ASCII art color                              |
| `--output=<file>`     | Saves output into a file                             |
| `--align=<type>`      | Aligns output (`left`, `right`, `center`, `justify`) |
| `--letters=<letters>` | Applies color to specific letters                    |
| `--reverse=<file>`    | Converts ASCII art back to text                      |

# Future Improvements

Possible enhancements that i am currently working on:

* Add colored ASCII output
* Support custom fonts
* Export output to files
* Add web interface
* Improve CLI argument parsing

---

# Author

Michael BAG

---