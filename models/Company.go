package models

type Company struct {

	// list you table columns here
	CompanyID   int    `json:"company_id" gorm:"primaryKey"`
	CompanyName string `json:"company_name"`
}
