package main

import "fmt"

type Meal struct {
	cost       int
	isTakeOut  bool
	mainCourse string
	drink      string
}

func NewMeal() *Meal {
	return &Meal{}
}

func (m *Meal) GetCost() int {
	return m.cost
}

func (m *Meal) GetIsTakeOut() bool {
	return m.isTakeOut
}

func (m *Meal) GetMainCourse() string {
	return m.mainCourse
}

func (m *Meal) GetDrink() string {
	return m.drink
}

func (m *Meal) SetCost(cost int) {
	m.cost = cost
}

func (m *Meal) SetIsTakeOut(takeout bool) {
	m.isTakeOut = takeout
}

func (m *Meal) SetMainCourse(course string) {
	m.mainCourse = course
}

func (m *Meal) SetDrink(drink string) {
	m.drink = drink
}

type MealBuilder struct {
	cost       int
	isTakeOut  bool
	mainCourse string
	drink      string
}

func NewMealBuilder() *MealBuilder {
	return &MealBuilder{}
}

func (mb *MealBuilder) AddCost(cost int) *MealBuilder {
	mb.cost = cost
	return mb
}

func (mb *MealBuilder) AddIsTakeOut(takeOut bool) *MealBuilder {
	mb.isTakeOut = takeOut
	return mb
}

func (mb *MealBuilder) AddMainCourse(course string) *MealBuilder {
	mb.mainCourse = course
	return mb
}

func (mb *MealBuilder) AddDrink(drink string) *MealBuilder {
	mb.drink = drink
	return mb
}

func (mb *MealBuilder) Build() *Meal {
	meal := NewMeal()
	meal.SetCost(mb.cost)
	meal.SetDrink(mb.drink)
	meal.SetIsTakeOut(mb.isTakeOut)
	meal.SetMainCourse(mb.mainCourse)
	return meal
}

type Director struct {
	mealBuilder *MealBuilder
}

func NewDirector(builder *MealBuilder) *Director {
	return &Director{mealBuilder: builder}
}

func (d *Director) BuildChickenComboMeal() *Meal {
	return d.mealBuilder.AddCost(45).AddIsTakeOut(true).AddMainCourse("Chicken").AddDrink("Zero Coke").Build()
}

func main() {
	builder := NewMealBuilder()
	director := NewDirector(builder)

	chickenComboMeal := director.BuildChickenComboMeal()

	fmt.Print(chickenComboMeal)
	fmt.Print("\n")
}
