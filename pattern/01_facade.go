package main

import "fmt"

type Subsys1 struct{}

func (*Subsys1) operation() {
	fmt.Println("Subsys1 запуск")
}

type Subsys2 struct{}

func (*Subsys2) operation() {
	fmt.Println("Subsys2 запуск")
}

type Facade struct {
	Subsys1
	Subsys2
}

func NewFacade() *Facade {
	return &Facade{
		Subsys1: Subsys1{},
		Subsys2: Subsys2{},
	}

}

func (f *Facade) Facadeoper() {
	fmt.Println("Запуск всех операций")
	f.Subsys1.operation()
	f.Subsys2.operation()

}

func main() {
	facade := NewFacade()
	facade.Facadeoper()
}


