package main

import "fmt"

type SquareHole struct {
	sideLength float64
}

func NewSquareHole(length float64) *SquareHole {
	return &SquareHole{
		sideLength: length,
	}
}

func (sh *SquareHole) CanFit(square ISquare) bool {
	return sh.sideLength >= square.GetSideLength()
}

type ISquare interface {
	GetSideLength() float64
}

type Square struct {
	sideLength float64
}

func NewSquare(length float64) *Square {
	return &Square{
		sideLength: length,
	}
}

func (s *Square) GetSideLength() float64 {
	return s.sideLength
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

type CircleAdapter struct {
	circle *Circle
}

func NewCircleAdapter(circle *Circle) *CircleAdapter {
	return &CircleAdapter{circle: circle}
}

func (ca *CircleAdapter) GetSideLength() float64 {
	return ca.circle.GetRadius()
}

type Stick struct {
	length float64
}

func NewStick(length float64) *Stick {
	return &Stick{
		length: length,
	}
}

func (st *Stick) GetStickLength() float64 {
	return st.length
}

type StickAdapter struct {
	stick *Stick
}

func NewStickAdapter(stick *Stick) *StickAdapter {
	return &StickAdapter{stick: stick}
}

func (sa *StickAdapter) GetSideLength() float64 {
	return sa.stick.GetStickLength()
}

func main() {
	squareHole := NewSquareHole(50)
	square := NewSquare(49)

	fmt.Println(squareHole.CanFit(square))

	circle := NewCircle(45)
	circleAdapter := NewCircleAdapter(circle)

	fmt.Println(squareHole.CanFit(circleAdapter))

	stick := NewStick(50)
	stickAdapter := NewStickAdapter(stick)

	fmt.Println(squareHole.CanFit(stickAdapter))
}
