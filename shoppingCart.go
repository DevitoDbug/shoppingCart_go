package main

type Item struct {
	itemId   string
	name     string
	price    float64
	quantity int
	deleted  bool
}
type ShoppingCart struct {
	items []Item
}

func (s ShoppingCart) GetAllItems() []Item {
	return s.items
}

func (s ShoppingCart) AddItemsToCart() {
	//TODO implement me
	panic("implement me")
}

func (s ShoppingCart) RemoveItem() {
	//TODO implement me
	panic("implement me")
}

func (s ShoppingCart) RestoreItem() {
	//TODO implement me
	panic("implement me")
}

func (s ShoppingCart) ForgetItem() {
	//TODO implement me
	panic("implement me")
}

func (s ShoppingCart) TotalPrice(totalPrice float64) {
	//TODO implement me
	panic("implement me")
}

func (s ShoppingCart) NumberOfItems() {
	//TODO implement me
	panic("implement me")
}

func (s ShoppingCart) CreateItem(name string, price float64, quantity int, deleted bool) *Item {
	item := &Item{
		itemId:   "something",
		name:     name,
		price:    price,
		quantity: quantity,
		deleted:  deleted,
	}
	return item
}
