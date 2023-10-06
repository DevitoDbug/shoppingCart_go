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
	NumberOfItems()
	CreateItem()
}

func main() {
	fmt.Println("Hello world!")
}
