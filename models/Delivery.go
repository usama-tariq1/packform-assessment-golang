package models

type Delivery struct {
	ID                int        `json:"id" gorm:"primaryKey"`
	OrderItemID       int        `json:"order_item_id"`
	OrderItem         *OrderItem `json:"order_item" gorm:"references:ID;foreignKey:OrderItemID" `
	DeliveredQuantity int        `json:"delivered_quantity"`
}
