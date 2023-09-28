package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

import (
	"fmt"
)

// Заказ представляет заказ, который нужно обработать.
type Order struct {
	ID           int
	TotalPrice   float64
	IsAuthorized bool
}

// Интерфейс обработчика заказов.
type OrderHandler interface {
	HandleOrder(order *Order)
	SetNextHandler(handler OrderHandler)
}

// Базовая реализация обработчика заказов.
type BaseOrderHandler struct {
	nextHandler OrderHandler
}

func (b *BaseOrderHandler) SetNextHandler(handler OrderHandler) {
	b.nextHandler = handler
}

// Конкретный обработчик заказов с проверкой на авторизацию.
type AuthorizationHandler struct {
	BaseOrderHandler
}

func (ah *AuthorizationHandler) HandleOrder(order *Order) {
	if order.IsAuthorized {
		fmt.Printf("Заказ %d: Авторизация пройдена.\n", order.ID)
		if ah.nextHandler != nil {
			ah.nextHandler.HandleOrder(order)
		}
	} else {
		fmt.Printf("Заказ %d: Ошибка авторизации. Заказ отклонен.\n", order.ID)
	}
}

// Конкретный обработчик заказов для вычисления стоимости.
type PriceCalculationHandler struct {
	BaseOrderHandler
}

func (pch *PriceCalculationHandler) HandleOrder(order *Order) {
	fmt.Printf("Заказ %d: Расчет стоимости заказа: %.2f\n", order.ID, order.TotalPrice)
	if pch.nextHandler != nil {
		pch.nextHandler.HandleOrder(order)
	}
}

func main() {
	authorizationHandler := &AuthorizationHandler{}
	priceCalculationHandler := &PriceCalculationHandler{}

	authorizationHandler.SetNextHandler(priceCalculationHandler)

	order1 := &Order{ID: 1, TotalPrice: 100.0, IsAuthorized: true}
	order2 := &Order{ID: 2, TotalPrice: 200.0, IsAuthorized: false}

	authorizationHandler.HandleOrder(order1)
	authorizationHandler.HandleOrder(order2)
}
