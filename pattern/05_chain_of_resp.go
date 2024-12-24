package main

import "fmt"

// Handler interface
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler struct
type BaseHandler struct {
	nextHandler Handler
}

// SetNext method for BaseHandler
func (b *BaseHandler) SetNext(handler Handler) {
	b.nextHandler = handler
}

// ConcreteHandler struct
type ConcreteHandler struct {
	BaseHandler
}

// Handle method for ConcreteHandler
func (c *ConcreteHandler) Handle(request string) {
	fmt.Println("ConcreteHandler handling request:", request)
	if c.nextHandler != nil {
		c.nextHandler.Handle(request)
	}
}

func main() {
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	handler1.Handle("Request")
}


/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
