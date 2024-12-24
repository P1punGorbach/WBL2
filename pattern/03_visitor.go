package main

import "fmt"

type Element interface {
	Accept(visitor Visitor)
}

type ElementA struct{}

func (cea *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(cea) // Передает себя посетителю
}

type ElementB struct{}

func (ceb *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(ceb)
}

type Visitor interface {
	VisitElementA(*ElementA)
	VisitElementB(*ElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitElementA(el *ElementA) {
	fmt.Println("Element A")
}
func (v *ConcreteVisitor) VisitElementB(el *ElementB) {
	fmt.Println("Element B")
}

type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) Attach(el Element) {
	os.elements = append(os.elements, el)
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, el := range os.elements {
		el.Accept(visitor)
	}
}

func main() {
	ObjectStructure := &ObjectStructure{}

	ObjectStructure.Attach(&ElementA{})
	ObjectStructure.Attach(&ElementB{})

	visitor := &ConcreteVisitor{}

	ObjectStructure.Accept(visitor)
}

