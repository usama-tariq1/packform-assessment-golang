package models

type OrderItem struct {
	// list you table columns here
	ID           int     `json:"id" gorm:"primaryKey"`
	OrderID      int     `json:"order_id"`
	Order        Order   `json:"order" gorm:"references:ID;foreignKey:OrderID"`
	PricePerUnit float64 `json:"price_per_unit"`
	Quantity     int     `json:"quantity"`
	Product      string  `json:"product"`
}
