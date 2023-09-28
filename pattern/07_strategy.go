package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

import (
	"fmt"
	"sort"
)

// Интерфейс Стратегии сортировки
type SortingStrategy interface {
	Sort([]int)
}

// Конкретная стратегия - сортировка пузырьком
type BubbleSortStrategy struct{}

func (bs *BubbleSortStrategy) Sort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// Конкретная стратегия - быстрая сортировка
type QuickSortStrategy struct{}

func (qs *QuickSortStrategy) Sort(data []int) {
	sort.Ints(data)
}

// Контекст, который использует выбранную стратегию сортировки
type Sorter struct {
	strategy SortingStrategy
}

func (s *Sorter) SetStrategy(strategy SortingStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(data []int) {
	s.strategy.Sort(data)
}

func main() {
	data := []int{9, 2, 5, 1, 6}

	sorter := &Sorter{}
	sorter.SetStrategy(&BubbleSortStrategy{})
	sorter.Sort(data)
	fmt.Println("Сортировка пузырьком:", data)

	sorter.SetStrategy(&QuickSortStrategy{})
	sorter.Sort(data)
	fmt.Println("Быстрая сортировка:", data)
}
