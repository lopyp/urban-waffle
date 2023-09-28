package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

import "fmt"

// Letter представляет письмо.
type Letter struct {
	Sender    string
	Recipient string
	Subject   string
	Body      string
}

// LetterBuilder - интерфейс для строителя письма.
type LetterBuilder interface {
	SetSender(sender string) LetterBuilder
	SetRecipient(recipient string) LetterBuilder
	SetSubject(subject string) LetterBuilder
	SetBody(body string) LetterBuilder
	Build() Letter
}

// ConcreteLetterBuilder - конкретная реализация строителя письма.
type ConcreteLetterBuilder struct {
	letter Letter
}

func NewLetterBuilder() LetterBuilder {
	return &ConcreteLetterBuilder{}
}

func (b *ConcreteLetterBuilder) SetSender(sender string) LetterBuilder {
	b.letter.Sender = sender
	return b
}

func (b *ConcreteLetterBuilder) SetRecipient(recipient string) LetterBuilder {
	b.letter.Recipient = recipient
	return b
}

func (b *ConcreteLetterBuilder) SetSubject(subject string) LetterBuilder {
	b.letter.Subject = subject
	return b
}

func (b *ConcreteLetterBuilder) SetBody(body string) LetterBuilder {
	b.letter.Body = body
	return b
}

func (b *ConcreteLetterBuilder) Build() Letter {
	return b.letter
}

func main() {
	letterBuilder := NewLetterBuilder()

	letter := letterBuilder.
		SetSender("john@example.com").
		SetRecipient("jane@example.com").
		SetSubject("Важное сообщение").
		SetBody("Привет, Джейн! Это важное сообщение от Джона.").
		Build()

	fmt.Printf("Письмо:\nОтправитель: %s\nПолучатель: %s\nТема: %s\nТекст:\n%s\n",
		letter.Sender, letter.Recipient, letter.Subject, letter.Body)
}
