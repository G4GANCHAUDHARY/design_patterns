package main

import "fmt"

type People struct {
	name      string
	age       int
	isMarried bool
}

func NewPeople(name string, age int, isMarried bool) *People {
	return &People{
		name:      name,
		age:       age,
		isMarried: isMarried,
	}
}

type PeopleCounter struct {
	Filter Filter
}

func NewPeopleCounter() *PeopleCounter {
	return &PeopleCounter{}
}

func (pc *PeopleCounter) SetFilter(filter Filter) {
	pc.Filter = filter
}

func (pc *PeopleCounter) Count(allPeeps []*People) int {
	return pc.Filter.GetCount(allPeeps)
}

type Filter interface {
	GetCount(allPeeps []*People) int
}

type AdultFilter struct {
}

func NewAdultFilter() *AdultFilter {
	return &AdultFilter{}
}

func (af *AdultFilter) GetCount(allPeeps []*People) (res int) {
	for _, p := range allPeeps {
		if p.age >= 18 {
			res += 1
		}
	}
	return res
}

type MarriedFilter struct {
}

func NewMarriedFilter() *MarriedFilter {
	return &MarriedFilter{}
}

func (mf *MarriedFilter) GetCount(allPeeps []*People) (res int) {
	for _, p := range allPeeps {
		if p.isMarried {
			res += 1
		}
	}
	return res
}

type SeniorFilter struct {
}

func NewSeniorFilter() *SeniorFilter {
	return &SeniorFilter{}
}

func (sf *SeniorFilter) GetCount(allPeeps []*People) (res int) {
	for _, p := range allPeeps {
		if p.age >= 65 {
			res += 1
		}
	}
	return res
}

func main() {
	allPeeps := []*People{NewPeople("Gagan", 24, false), NewPeople("Prerit", 25, false), NewPeople("Athrva", 26, true)}
	peepCounter := NewPeopleCounter()

	peepCounter.SetFilter(NewAdultFilter())
	fmt.Println(peepCounter.Count(allPeeps))

	peepCounter.SetFilter(NewSeniorFilter())
	fmt.Println(peepCounter.Count(allPeeps))

	peepCounter.SetFilter(NewMarriedFilter())
	fmt.Println(peepCounter.Count(allPeeps))
}
