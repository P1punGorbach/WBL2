package main

import "fmt"

type House struct { //Объект который нужно построить
	roof string
	wall string
	door string
}
type Builder interface { // Определение методов для пошагового построения объекта
	BuildWall()
	BuildRoof()
	BuildDoor()
	GetHouse() House
}

type ConcreteBuilder struct { //Строитель, реализует методы интерфейса
	house House
}

func (b *ConcreteBuilder) BuildWall() {
	b.house.wall = "Каменная стена"
}

func (b *ConcreteBuilder) BuildDoor() {
	b.house.door = "Деревянная стена"
}

func (b *ConcreteBuilder) BuildRoof() {
	b.house.roof = "Черепица стена"
}

func (b *ConcreteBuilder) GetHouse() House { // Метод получения готового объекта
	return b.house
}

type Director struct { // Директор, управляет процессом сборки объекта
	builder Builder
}

func (d *Director) SetBuilder(b Builder) { // Метод назначение строителя директору
	d.builder = b
}

func (d *Director) Construct() { // Вызов методов строителя в определенной последовательности
	d.builder.BuildDoor()
	d.builder.BuildRoof()
	d.builder.BuildWall()
}

func main() {
	builder := &ConcreteBuilder{}
	director := &Director{}

	director.SetBuilder(builder)
	director.Construct()

	house := builder.GetHouse()
	fmt.Println("Дом состоит из:")
	fmt.Println("Стены: ", house.wall)
	fmt.Println("Крыша: ", house.roof)
	fmt.Println("Дверь: ", house.door)
}


