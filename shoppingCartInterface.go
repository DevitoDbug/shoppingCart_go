package main

import (
	"fmt"
)

type ShoppingCartInterface interface {
	GetAllItems() []Item
	AddItemsToCart()
	RemoveItem()
	RestoreItem()
	ForgetItem()
	TotalPrice(totalPrice float64)
	NumberOfItems() int
	CreateItem(name string, price float64, quantity int, deleted bool) *Item
}

func main() {
	fmt.Println("Hello world!")
}
