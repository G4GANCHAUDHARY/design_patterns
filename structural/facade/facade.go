package main

import "fmt"

type House struct {
	window    int
	door      int
	furniture int
}

func NewHouse(w int, d int, f int) *House {
	return &House{
		window:    w,
		door:      d,
		furniture: f,
	}
}

type HouseFacade struct {
}

func NewHouseFacade() *HouseFacade {
	return &HouseFacade{}
}

func (hf *HouseFacade) BuildSmallHouse() *House {
	return NewHouse(4, 2, 0)
}

func (hf *HouseFacade) BuildMediumHouse() *House {
	return NewHouse(8, 4, 2)
}

func (hf *HouseFacade) BuildBigHouse() *House {
	return NewHouse(16, 12, 6)
}

func main() {
	houseFacade := NewHouseFacade()
	fmt.Println(houseFacade.BuildSmallHouse())
	fmt.Println(houseFacade.BuildMediumHouse())
	fmt.Println(houseFacade.BuildBigHouse())
}
