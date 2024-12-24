package main

import "fmt"

type Strategy interface {
	DoOperation(a, b int) int
}

type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) DoOperation(a, b int) int {
	return a + b
}

type ConcreteStrategyMultiply struct{}

func (s *ConcreteStrategyMultiply) DoOperation(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.DoOperation(a, b)
}

func main() {

	context := &Context{}

	context.SetStrategy(&ConcreteStrategyAdd{})
	fmt.Printf("10 + 5 = %d\n", context.ExecuteStrategy(10, 5))

	context.SetStrategy(&ConcreteStrategyMultiply{})
	fmt.Printf("10 * 5 = %d\n", context.ExecuteStrategy(10, 5))
}

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
