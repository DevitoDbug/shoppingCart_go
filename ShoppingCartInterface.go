package ShoppingCartInterface

type ShoppingCart interface {
	GetAllItems()
	AddItemsToCart()
	RemoveItem()
	RestoreItem()
	ForgetItem()
	TotalPrice(totalPrice float64)
	NumberOfItems()
}

func main() {

}
