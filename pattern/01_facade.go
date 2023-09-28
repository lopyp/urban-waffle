package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

import "fmt"

// Подсистема 1
type Subsystem1 struct {
}

func (s *Subsystem1) Operation1() {
	fmt.Println("Subsystem1: Operation1")
}

func (s *Subsystem1) Operation2() {
	fmt.Println("Subsystem1: Operation2")
}

// Подсистема 2
type Subsystem2 struct {
}

func (s *Subsystem2) Operation3() {
	fmt.Println("Subsystem2: Operation3")
}

func (s *Subsystem2) Operation4() {
	fmt.Println("Subsystem2: Operation4")
}

// Фасад
type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

func (f *Facade) Operation() {
	fmt.Println("Facade: операция начата")
	f.subsystem1.Operation1()
	f.subsystem1.Operation2()
	f.subsystem2.Operation3()
	f.subsystem2.Operation4()
	fmt.Println("Facade: операция завершена")
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
