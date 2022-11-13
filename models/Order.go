package models

import "time"

type Order struct {
	// list you table columns here
	ID         int       `json:"id" gorm:"primaryKey"`
	OrderName  string    `json:"order"`
	CustomerID string    `json:"customer_id"`
	Customer   Customer  `gorm:"references:UserID;foreignKey:CustomerID" json:"customer,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
