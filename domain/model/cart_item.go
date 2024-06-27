package model

type CartItem struct {
	Id          string `json:"id"`
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int32  `json:"price"`
}
