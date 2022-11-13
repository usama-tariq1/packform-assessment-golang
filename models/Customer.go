package models

import "github.com/lib/pq"

type Customer struct {
	// list you table columns here
	UserID      string         `json:"user_id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Password    string         `json:"-,omitempty"`
	CompanyID   int            `json:"company_id"`
	Company     Company        `gorm:"references:CompanyID;foreignKey:CompanyID" json:"company"`
	Login       string         `json:"login"`
	CreditCards pq.StringArray `json:"-,omitempty" gorm:"type:text[]"`
}
