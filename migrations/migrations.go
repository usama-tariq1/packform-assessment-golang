package migrations

import (
	"github.com/usama-tariq1/leet-gin/helper"
	"github.com/usama-tariq1/leet-gin/models"
	"gorm.io/gorm"
)

// var DB *gorm.DB
var console helper.Console
var customer models.Customer
var company models.Company
var order models.Order
var orderItem models.OrderItem
var delivery models.Delivery

// list models here
func Migrate(DB *gorm.DB) {

	DB.AutoMigrate(company)
	DB.AutoMigrate(customer)
	DB.AutoMigrate(order)
	DB.AutoMigrate(orderItem)
	DB.AutoMigrate(delivery)

}
