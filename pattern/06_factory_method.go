package main

import "fmt"

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	Name  string
	Power int
}

func (g *Gun) setName(name string) {
	g.Name = name
}

func (g *Gun) setPower(power int) {
	g.Power = power
}

func (g *Gun) getName() string {
	return g.Name
}

func (g *Gun) getPower() int {
	return g.Power
}

type AK47 struct {
	Gun
}

func NewAK47() iGun {
	return &AK47{
		Gun{
			"AK47",
			100,
		},
	}
}

type Musket struct {
	Gun
}

func NewMusket() iGun {
	return &Musket{
		Gun: Gun{
			"Musket",
			500,
		},
	}
}

func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return NewAK47(), nil
	case "musket":
		return NewMusket(), nil
	default:
		return nil, fmt.Errorf("Unknown gun")
	}
}

func main() {
	g1, err := getGun("ak47")
	fmt.Println(g1.getName())

	g2, err := getGun("musket")
	fmt.Println(g2.getName())

	_, err = getGun("something")
	if err != nil {
		fmt.Println(err.Error())
	}
}

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
