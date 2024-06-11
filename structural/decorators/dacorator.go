package main

import "fmt"

type Coffee interface {
	GetCost() float64
	GetTyp() string
}

type SimpleCoffee struct {
	typ string
}

func NewSimpleCoffee(typ string) *SimpleCoffee {
	return &SimpleCoffee{typ: typ}
}

func (sc *SimpleCoffee) GetCost() float64 {
	return 5.5
}

func (sc *SimpleCoffee) GetTyp() string {
	return sc.typ
}

type MilkDecorator struct {
	typ    string
	coffee Coffee
}

func NewMilkDecorator(typ string, coffee Coffee) *MilkDecorator {
	return &MilkDecorator{typ: typ, coffee: coffee}
}

func (c *MilkDecorator) GetCost() float64 {
	return c.coffee.GetCost() + 20.05
}

func (c *MilkDecorator) GetTyp() string {
	return c.typ
}

type CaramelDecorator struct {
	typ    string
	coffee Coffee
}

func NewCaramelDecorator(typ string, coffee Coffee) *CaramelDecorator {
	return &CaramelDecorator{typ: typ, coffee: coffee}
}

func (c *CaramelDecorator) GetCost() float64 {
	return c.coffee.GetCost() + 10.05
}

func (c *CaramelDecorator) GetTyp() string {
	return c.typ
}

func main() {
	simple := NewSimpleCoffee("Simple Coffee")
	fmt.Println(simple.GetTyp())
	fmt.Println(simple.GetCost())

	milkDecoratedCoffee := NewMilkDecorator("Milk Decorated Coffee", simple)
	fmt.Println(simple.GetTyp(), milkDecoratedCoffee.GetTyp())
	fmt.Println(milkDecoratedCoffee.GetCost())

	caramelDecoratedCoffee := NewCaramelDecorator("Caramel Decorated Coffee", milkDecoratedCoffee)
	fmt.Println(simple.GetTyp(), milkDecoratedCoffee.GetTyp(), caramelDecoratedCoffee.GetTyp())
	fmt.Println(caramelDecoratedCoffee.GetCost())
}
