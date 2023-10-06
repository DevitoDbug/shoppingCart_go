package main

import (
	"fmt"
)

type ShoppingCartInterface interface {
	GetAllItems() []Item
	AddItemsToCart(item Item)
	RemoveItem(itemId string)
	RestoreItem(itemId string)
	ForgetItem(itemId string)
	TotalPrice() (totalPrice float64)
	NumberOfItems() (numberOfItems int)
	CreateItem(name string, price float64, quantity int, deleted bool) *Item
}

func main() {
	fmt.Println("Hello world!")
}
