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

func (s ShoppingCart) AddItemsToCart(item Item) {
	s.items = append(s.items, item)
}

func (s ShoppingCart) RemoveItem(itemId string) {
	items := s.items
	for _, item := range items {
		if item.itemId == itemId {
			item.deleted = true
			return
		}
	}
}

func (s ShoppingCart) RestoreItem(itemId string) {
	items := s.items
	for _, item := range items {
		if item.itemId == itemId {
			item.deleted = false
			return
		}
	}
}

func (s ShoppingCart) ForgetItem(itemId string) {
	items := s.items
	for index, item := range items {
		if item.itemId == itemId {
			items = append(items[:index], items[index+1:]...)
			s.items = items
			return
		}
	}
}

func (s ShoppingCart) TotalPrice() (totalPrice float64) {
	var TotalPrice float64
	items := s.items
	for _, item := range items {
		TotalPrice += float64(item.quantity) * item.price
	}
	return TotalPrice
}

func (s ShoppingCart) NumberOfItems() int {
	items := s.items
	var numberOfItems int
	for _, item := range items {
		numberOfItems += item.quantity
	}
	return numberOfItems
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
