package main

import "fmt"

type Customer interface {
	UpdateCustomers(count int)
	GetName() string
	GetNotificationCounts() int
}

type BargainHunters struct {
	name               string
	notificationCounts int
}

func NewBargainHunters(name string) *BargainHunters {
	return &BargainHunters{name: name, notificationCounts: 0}
}

func (bh *BargainHunters) UpdateCustomers(count int) {
	fmt.Println("Bargain Hunters has been informed about the item stock. New Stock : ", count)
	bh.notificationCounts++
	return
}

func (bh *BargainHunters) GetNotificationCounts() int {
	return bh.notificationCounts
}

func (bh *BargainHunters) GetName() string {
	return bh.name
}

type WindowShoppers struct {
	name               string
	notificationCounts int
}

func NewWindowShoppers(name string) *WindowShoppers {
	return &WindowShoppers{name: name, notificationCounts: 0}
}

func (ws *WindowShoppers) UpdateCustomers(count int) {
	fmt.Println("Window Shoppers has been informed about the item stock. New Stock : ", count)
	ws.notificationCounts++
	return
}

func (ws *WindowShoppers) GetNotificationCounts() int {
	return ws.notificationCounts
}

func (ws *WindowShoppers) GetName() string {
	return ws.name
}

type OnlineStoreItem struct {
	product     string
	stockCount  int
	subscribers []Customer
}

func NewOnlineStoreItem(product string, stock int) *OnlineStoreItem {
	return &OnlineStoreItem{
		product:     product,
		stockCount:  stock,
		subscribers: make([]Customer, 0),
	}
}

func (si *OnlineStoreItem) UpdateStock(count int) {
	if si.stockCount == 0 && count != 0 {
		for _, c := range si.subscribers {
			c.UpdateCustomers(count)
		}
	}
	si.stockCount = count
}

func (si *OnlineStoreItem) Unsubscribe(customer Customer) {
	for i, c := range si.subscribers {
		if c.GetName() == customer.GetName() {
			si.subscribers = append(si.subscribers[:i], si.subscribers[i+1:]...)
			return
		}
	}
}

func (si *OnlineStoreItem) Subscribe(customer Customer) {
	si.subscribers = append(si.subscribers, customer)
}

func main() {
	onlineStoreItem := NewOnlineStoreItem("book", 0)

	bargainHunters := NewBargainHunters("Bargain Hunters")

	windowShoppers := NewWindowShoppers("Window Shoppers")

	onlineStoreItem.Subscribe(bargainHunters)
	onlineStoreItem.Subscribe(windowShoppers)

	onlineStoreItem.UpdateStock(5)

	onlineStoreItem.Unsubscribe(bargainHunters)

	onlineStoreItem.UpdateStock(0)
	onlineStoreItem.UpdateStock(3)
	onlineStoreItem.UpdateStock(2)

	fmt.Println(bargainHunters.GetNotificationCounts())
	fmt.Println(windowShoppers.GetNotificationCounts())
}
