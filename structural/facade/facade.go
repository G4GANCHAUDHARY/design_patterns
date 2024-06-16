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

func (h *House) SetWindows(w int) {
	h.window = w
}

func (h *House) SetDoors(d int) {
	h.window = d
}

func (h *House) SetFurniture(f int) {
	h.window = f
}

type HouseFacade struct {
	House *House
}

func NewHouseFacade(house *House) *HouseFacade {
	return &HouseFacade{House: house}
}

func (hf *HouseFacade) SetSmallHouse() {
	hf.House.SetDoors(2)
	hf.House.SetWindows(2)
	hf.House.SetFurniture(2)
}

func (hf *HouseFacade) SetMediumHouse() {
	hf.House.SetDoors(4)
	hf.House.SetWindows(4)
	hf.House.SetFurniture(4)
}

func (hf *HouseFacade) SetBigHouse() {
	hf.House.SetDoors(6)
	hf.House.SetWindows(6)
	hf.House.SetFurniture(6)
}

func main() {
	house := NewHouse(0, 0, 0)
	houseFacade := NewHouseFacade(house)
	houseFacade.SetSmallHouse()
	fmt.Println(house)
	houseFacade.SetMediumHouse()
	fmt.Println(house)
	houseFacade.SetBigHouse()
	fmt.Println(house)
}
