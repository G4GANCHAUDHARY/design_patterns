package main

import "fmt"

type Watches interface {
	GetName() string
	GetPrice() int
}

type Titan struct {
	name  string
	price int
}

func NewTitanWatch(name string, price int) *Titan {
	return &Titan{
		name:  name,
		price: price,
	}
}

func (t *Titan) GetName() string {
	return t.name
}

func (t *Titan) GetPrice() int {
	return t.price
}

type Rolex struct {
	name  string
	price int
}

func NewRolexWatch(name string, price int) *Rolex {
	return &Rolex{
		name:  name,
		price: price,
	}
}

func (r *Rolex) GetName() string {
	return r.name
}

func (r *Rolex) GetPrice() int {
	return r.price
}

type WatchFactory struct{}

func NewWatchFactory() *WatchFactory {
	return &WatchFactory{}
}

func (w *WatchFactory) GetNewWatch(typ string) (string, int) {
	var watch Watches
	if typ == "Titan" {
		watch = NewTitanWatch(typ, 1000)
	} else if typ == "Rolex" {
		watch = NewRolexWatch(typ, 100000)
	}

	return watch.GetName(), watch.GetPrice()
}

func main() {
	factory := NewWatchFactory()
	watch, price := factory.GetNewWatch("Titan")
	fmt.Printf("Client want %v watch priced at %v !\n", watch, price)

	watch, price = factory.GetNewWatch("Rolex")
	fmt.Printf("Client want %v watch priced at %v !\n", watch, price)
}
