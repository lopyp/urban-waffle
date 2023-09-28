package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	AfterLines   int
	BeforeLines  int
	ContextLines int
	CountOnly    bool
	IgnoreCase   bool
	InvertMatch  bool
	FixedString  bool
	LineNumber   bool
	SearchString string
	FilePath     string
}

func main() {
	config := parseCommandLine()

	input := openInputFile(config.FilePath)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	lineNumberCounter := 1
	matched := false

	for scanner.Scan() {
		line := scanner.Text()
		lineToMatch := line

		if config.IgnoreCase {
			lineToMatch = strings.ToLower(line)
			config.SearchString = strings.ToLower(config.SearchString)
		}

		if config.FixedString {
			matched = (lineToMatch == config.SearchString)
		} else {
			matched = strings.Contains(lineToMatch, config.SearchString)
		}

		if config.InvertMatch {
			matched = !matched
		}

		if matched {
			printMatchedLine(line, lineNumberCounter, config)
			matched = false
			if config.AfterLines > 0 {
				printFollowingLines(scanner, config.AfterLines, lineNumberCounter)
			}
		} else if config.BeforeLines > 0 {
			printLeadingLines(line, lineNumberCounter, config)
		}

		lineNumberCounter++
	}

	if config.CountOnly {
		fmt.Println("Matching lines:", lineNumberCounter-1)
	}
}

func parseCommandLine() Config {
	config := Config{}
	flag.IntVar(&config.AfterLines, "A", 0, "Print +N lines after each match")
	flag.IntVar(&config.BeforeLines, "B", 0, "Print +N lines before each match")
	flag.IntVar(&config.ContextLines, "C", 0, "Print ±N lines around each match")
	flag.BoolVar(&config.CountOnly, "c", false, "Count the number of matching lines")
	flag.BoolVar(&config.IgnoreCase, "i", false, "Ignore case when matching")
	flag.BoolVar(&config.InvertMatch, "v", false, "Invert the match (select non-matching lines)")
	flag.BoolVar(&config.FixedString, "F", false, "Match the fixed string (not a pattern)")
	flag.BoolVar(&config.LineNumber, "n", false, "Print line numbers")
	flag.Parse()
	config.SearchString = flag.Arg(0)
	config.FilePath = flag.Arg(1)
	return config
}

func openInputFile(filePath string) *os.File {
	if filePath != "" {
		input, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		return input
	}
	return os.Stdin
}

func printMatchedLine(line string, lineNumber int, config Config) {
	if config.LineNumber {
		fmt.Printf("%d:", lineNumber)
	}
	fmt.Println(line)
}

func printLeadingLines(line string, lineNumber int, config Config) {
	if lineNumber <= config.BeforeLines {
		printMatchedLine(line, lineNumber, config)
	}
}

func printFollowingLines(scanner *bufio.Scanner, numLines int, startLineNumber int) {
	for i := 1; i <= numLines; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			printMatchedLine(line, startLineNumber+i, Config{})
		} else {
			break
		}
	}
}
