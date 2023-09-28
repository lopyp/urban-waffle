package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

import "fmt"

// Интерфейс для элемента, который может быть посещен посетителем.
type Shape interface {
	Accept(Visitor)
}

// Конкретные элементы - геометрические фигуры.
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

// Интерфейс для посетителя.
type Visitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}

// Конкретный посетитель - вычисляет площадь фигур.
type AreaVisitor struct {
	Area float64
}

func (av *AreaVisitor) VisitCircle(c *Circle) {
	av.Area += 3.14 * c.Radius * c.Radius
}

func (av *AreaVisitor) VisitRectangle(r *Rectangle) {
	av.Area += r.Width * r.Height
}

func main() {
	circle := &Circle{Radius: 5.0}
	rectangle := &Rectangle{Width: 4.0, Height: 6.0}

	areaVisitor := &AreaVisitor{}

	circle.Accept(areaVisitor)
	rectangle.Accept(areaVisitor)

	fmt.Printf("Площадь круга: %.2f\n", areaVisitor.Area)
	fmt.Printf("Площадь прямоугольника: %.2f\n", areaVisitor.Area)
}
