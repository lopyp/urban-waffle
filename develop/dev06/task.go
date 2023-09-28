package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Определение флагов
	fieldNumbersStr := flag.String("f", "", "Fields to select (comma-separated)")
	delimiterStr := flag.String("d", "\t", "Delimiter")
	separatedOnly := flag.Bool("s", false, "Only output lines with the delimiter")
	flag.Parse()

	// Получение чисел полей для выбора
	fieldNumbers := parseFieldNumbers(*fieldNumbersStr)

	// Установка разделителя
	delimiter := *delimiterStr

	// Считывание входных данных со стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if *separatedOnly && !strings.Contains(line, delimiter) {
			continue
		}

		fields := strings.Split(line, delimiter)
		selectedFields := selectFields(fields, fieldNumbers)

		fmt.Println(strings.Join(selectedFields, delimiter))
	}
}

func parseFieldNumbers(fieldNumbersStr string) []int {
	if fieldNumbersStr == "" {
		return []int{}
	}

	fields := strings.Split(fieldNumbersStr, ",")
	fieldNumbers := make([]int, len(fields))

	for i, field := range fields {
		fieldNumbers[i] = parseInt(field)
	}

	return fieldNumbers
}

func parseInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing integer: %v\n", err)
		os.Exit(1)
	}
	return num
}

func selectFields(fields []string, fieldNumbers []int) []string {
	selectedFields := make([]string, len(fieldNumbers))

	for i, num := range fieldNumbers {
		if num >= 1 && num <= len(fields) {
			selectedFields[i] = fields[num-1]
		} else {
			selectedFields[i] = ""
		}
	}

	return selectedFields
}
