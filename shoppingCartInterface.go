package main

import (
	"fmt"
)

type ShoppingCart interface {
	GetAllItems()
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
