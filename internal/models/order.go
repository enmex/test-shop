package models

type Order struct {
	ID       int
	Products []*ProductInOrder
}

type ProductInOrder struct {
	Product
	Quantity int
}
