package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

import "fmt"

// Интерфейс Команды
type Command interface {
	Execute()
}

// Получатель - автомобиль
type Car struct {
	IsRunning bool
}

func (c *Car) Start() {
	c.IsRunning = true
	fmt.Println("Автомобиль заведен")
}

func (c *Car) Stop() {
	c.IsRunning = false
	fmt.Println("Автомобиль заглушен")
}

// Конкретная Команда для запуска автомобиля
type StartCommand struct {
	Car *Car
}

func (sc *StartCommand) Execute() {
	sc.Car.Start()
}

// Конкретная Команда для остановки автомобиля
type StopCommand struct {
	Car *Car
}

func (sc *StopCommand) Execute() {
	sc.Car.Stop()
}

// Инвокер - ключ зажигания автомобиля
type IgnitionKey struct {
	Command Command
}

func (ik *IgnitionKey) Turn() {
	ik.Command.Execute()
}

func main() {

	car := &Car{}

	startCommand := &StartCommand{Car: car}
	stopCommand := &StopCommand{Car: car}

	ignitionKey := &IgnitionKey{}

	ignitionKey.Command = startCommand
	ignitionKey.Turn()

	ignitionKey.Command = stopCommand
	ignitionKey.Turn()
}
