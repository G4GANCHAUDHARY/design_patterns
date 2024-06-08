package main

import "fmt"

type OgInterface interface {
	GetId() int
	GetName() string
	CloneOriginal() OgInterface
}

type Original1 struct {
	id   int
	name string
}

func NewOg1(id int, name string) *Original1 {
	return &Original1{id: id, name: name}
}

func (o *Original1) GetId() int {
	return o.id
}

func (o *Original1) GetName() string {
	return o.name
}

func (o *Original1) CloneOriginal() OgInterface {
	newId := o.GetId()
	newName := o.GetName()
	return &Original1{
		id:   newId,
		name: newName,
	}
}

type Original2 struct {
	id   int
	name string
}

func NewOg2(id int, name string) *Original2 {
	return &Original2{id: id, name: name}
}

func (o *Original2) GetId() int {
	return o.id
}

func (o *Original2) GetName() string {
	return o.name
}

func (o *Original2) CloneOriginal() OgInterface {
	newId := o.GetId()
	newName := o.GetName()
	return &Original2{
		id:   newId,
		name: newName,
	}
}

func main() {
	og1 := NewOg1(1, "og1")
	og2 := NewOg2(2, "og2")

	fmt.Print(og1)
	fmt.Print("\n")
	fmt.Print(og2)
	fmt.Print("\n")

	cloneOg1 := og1.CloneOriginal()
	cloneOg2 := og2.CloneOriginal()

	fmt.Print(cloneOg1)
	fmt.Print("\n")

	fmt.Print(cloneOg2)
	fmt.Print("\n")

}
