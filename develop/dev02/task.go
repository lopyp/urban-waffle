package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"strings"
	"unicode"
)

func unpackString(s string) (string, error) {
	var result strings.Builder
	escape := false
	num_escape := false

	for i := 0; i < len(s); i++ {
		if escape {
			if unicode.IsDigit(rune(s[i])) || s[i] == '\\' {
				result.WriteRune(rune(s[i]))
				num_escape = true
			} else {
				return "", errors.New("некорректная строка")
			}
			escape = false
		} else if s[i] == '\\' {
			if i == len(s)-1 {
				return "", errors.New("некорректная строка")
			}
			escape = true
		} else {
			if unicode.IsDigit(rune(s[i])) {
				count := int(s[i] - '0')
				if i == 0 || (unicode.IsDigit(rune(s[i-1])) && !num_escape) || count == 0 {
					return "", errors.New("некорректная строка")
				}
				result.WriteString(strings.Repeat(string(s[i-1]), count-1))
				num_escape = true
			} else if unicode.IsLetter(rune(s[i])) {
				result.WriteRune(rune(s[i]))
			} else {
				return "", errors.New("некорректная строка")
			}
		}
	}

	return result.String(), nil
}
