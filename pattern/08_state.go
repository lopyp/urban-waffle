package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

import (
	"fmt"
)

// Интерфейс состояния чая
type TeaState interface {
	DrinkTea()
}

// Конкретное состояние "Холодный чай"
type ColdTeaState struct{}

func (cts *ColdTeaState) DrinkTea() {
	fmt.Println("Вы пьете холодный чай. брр!")
}

// Конкретное состояние "Теплый чай"
type WarmTeaState struct{}

func (wts *WarmTeaState) DrinkTea() {
	fmt.Println("Вы пьете теплый чай. Приятно!")
}

// Конкретное состояние "Горячий чай"
type HotTeaState struct{}

func (hts *HotTeaState) DrinkTea() {
	fmt.Println("Вы пьете горячий чай. Осторожно!")
}

// Чашка чая
type TeaCup struct {
	state TeaState
}

func NewTeaCup() *TeaCup {
	return &TeaCup{state: &ColdTeaState{}}
}

func (tc *TeaCup) SetState(state TeaState) {
	tc.state = state
}

func (tc *TeaCup) DrinkTea() {
	tc.state.DrinkTea()
}

func main() {
	teaCup := NewTeaCup()

	teaCup.DrinkTea()

	teaCup.SetState(&WarmTeaState{})
	teaCup.DrinkTea()

	teaCup.SetState(&HotTeaState{})
	teaCup.DrinkTea()
}
