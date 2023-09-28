package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range words {
		// Приводим слово к нижнему регистру и сортируем его буквы
		word = strings.ToLower(word)
		wordRune := []rune(word)
		sort.Slice(wordRune, func(i, j int) bool { return wordRune[i] < wordRune[j] })
		sortedWord := string(wordRune)

		// Добавляем слово в соответствующее множество анаграмм
		if _, found := anagramSets[sortedWord]; !found {
			anagramSets[sortedWord] = []string{word}
		} else {
			anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
		}
	}

	// Удаляем множества из одного элемента
	for key, value := range anagramSets {
		if len(value) == 1 {
			delete(anagramSets, key)
		}
	}

	return anagramSets
}
