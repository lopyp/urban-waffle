package main

/*
=== Утилита sort ===



Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	SortColumn int
	Numeric    bool
	Reverse    bool
	Unique     bool
}

func main() {
	config := Config{}
	flag.IntVar(&config.SortColumn, "k", 0, "номер колонки для сортировки (начиная с 1)")
	flag.BoolVar(&config.Numeric, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&config.Reverse, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&config.Unique, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Использование: go run sort.go <имя файла>")
		return
	}

	fileName := flag.Arg(0)
	lines, err := readLinesFromFile(fileName)
	if err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		return
	}

	sort.SliceStable(lines, func(i, j int) bool {
		return compareLines(lines[i], lines[j], config.SortColumn-1, config.Numeric)
	})

	if config.Unique {
		seen := make(map[string]struct{})
		var uniqueLines []string
		for _, line := range lines {
			if _, ok := seen[line]; !ok {
				uniqueLines = append(uniqueLines, line)
				seen[line] = struct{}{}
			}
		}
		lines = uniqueLines
	}

	if config.Reverse {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func readLinesFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func splitColumns(line string) []string {
	var columns []string
	for _, col := range strings.Split(line, "\t") {
		columns = append(columns, col)
	}
	return columns
}

func compareLines(line1, line2 string, sortColumn int, numeric bool) bool {
	columns1 := splitColumns(line1)
	columns2 := splitColumns(line2)

	if sortColumn < 0 || sortColumn >= len(columns1) || sortColumn >= len(columns2) {
		return false
	}

	if numeric {
		val1, err1 := strconv.Atoi(columns1[sortColumn])
		val2, err2 := strconv.Atoi(columns2[sortColumn])
		if err1 == nil && err2 == nil {
			return val1 < val2
		}
	}

	return columns1[sortColumn] < columns2[sortColumn]
}
