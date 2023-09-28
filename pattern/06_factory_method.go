package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

import "fmt"

// Интерфейс Транспортного средства
type Transport interface {
	Drive()
}

// Конкретный тип Транспортного средства - Автомобиль
type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Едем на автомобиле")
}

// Конкретный тип Транспортного средства - Мотоцикл
type Motorcycle struct{}

func (m *Motorcycle) Drive() {
	fmt.Println("Едем на мотоцикле")
}

// Интерфейс Фабрики Транспортных средств
type TransportFactory interface {
	CreateTransport() Transport
}

// Конкретная Фабрика для создания Автомобилей
type CarFactory struct{}

func (cf CarFactory) CreateTransport() Transport {
	return &Car{}
}

// Конкретная Фабрика для создания Мотоциклов
type MotorcycleFactory struct{}

func (mf MotorcycleFactory) CreateTransport() Transport {
	return &Motorcycle{}
}

func main() {

	carFactory := CarFactory{}
	motorcycleFactory := MotorcycleFactory{}

	car := carFactory.CreateTransport()
	motorcycle := motorcycleFactory.CreateTransport()

	car.Drive()
	motorcycle.Drive()
}
