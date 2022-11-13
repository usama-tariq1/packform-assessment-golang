package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/usama-tariq1/leet-gin/models"
	"gorm.io/gorm"
)

type OrderController struct {
}

type DeliveryArrayResponse struct {
	ID              int     `json:"id"`
	OrderName       string  `json:"order_name"`
	ProductName     string  `json:"product_name"`
	CustomerCompany string  `json:"customer_company"`
	CustomerName    string  `json:"customer_name"`
	OrderDate       string  `json:"order_date"`
	DeliveredAmount float64 `json:"delivered_amount"`
	TotalAmount     float64 `json:"total_amount"`
}

type OrderResponse struct {
	Data  []DeliveryArrayResponse `json:"data"`
	Total int64                   `json:"total"`
}

var ()

func (OrderController OrderController) Index(app *gin.Context) {
	var deliveries []models.Delivery
	var count int64

	// query with ORM
	// Lamp().Scopes(Paginate(app)).Preload("OrderItem.Order.Customer.Company").Find(&deliveries)
	// There were some limitation with orm i was using the where and joins were not working in conjunction
	// therefor i used some manual filtering

	productFilter := app.DefaultQuery("search_query", "")
	orderSort := app.DefaultQuery("sort_by", "")
	startTime := app.DefaultQuery("start_time", "")
	endTime := app.DefaultQuery("end_time", "")
	rangeFilter := ""

	if productFilter != "" || orderSort != "" || startTime != "" {
		if productFilter != "" {
			productFilter = fmt.Sprintf("where order_items.product like '%%%s%%' or orders.order_name like '%%%s%%'", productFilter, productFilter)
		}

		if orderSort != "" {
			orderSort = fmt.Sprintf("order by %s", orderSort)
		}

		if startTime != "" {
			if productFilter != "" {
				rangeFilter = fmt.Sprintf(" and orders.created_at between '%s' and '%s' ", startTime, endTime)

			} else {
				rangeFilter = fmt.Sprintf("where orders.created_at between '%s' and '%s' ", startTime, endTime)

			}
		}

		query := fmt.Sprintf(`left join  order_items on order_items.id= deliveries.order_item_id 
		left join orders on orders.id= order_items.order_id %s %s %s`, productFilter, rangeFilter, orderSort)
		Lamp().Joins(query).
			Preload("OrderItem.Order.Customer.Company").Scopes(Paginate(app)).Find(&deliveries).Count(&count)
	} else {

		Lamp().Scopes(Paginate(app)).Preload("OrderItem.Order.Customer.Company").Find(&deliveries).Count(&count)

	}

	// Lamp().Model(&deliveryModel).Count(&count)

	var response []DeliveryArrayResponse
	for _, delivery := range deliveries {
		var deliveryResponse DeliveryArrayResponse
		deliveryResponse = DeliveryArrayResponse{
			ID:              delivery.ID,
			OrderName:       delivery.OrderItem.Order.OrderName,
			ProductName:     delivery.OrderItem.Product,
			CustomerCompany: delivery.OrderItem.Order.Customer.Company.CompanyName,
			CustomerName:    delivery.OrderItem.Order.Customer.Name,
			OrderDate:       delivery.OrderItem.Order.CreatedAt.UTC().String(),
			DeliveredAmount: float64(delivery.DeliveredQuantity) * delivery.OrderItem.PricePerUnit,
			TotalAmount:     float64(delivery.OrderItem.Quantity) * delivery.OrderItem.PricePerUnit,
		}
		response = append(response, deliveryResponse)
	}

	app.IndentedJSON(http.StatusOK, OrderResponse{
		Total: count,
		Data:  response,
	})
}

func (OrderController OrderController) Create(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Create Called")

}

func (OrderController OrderController) Read(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Read Called")

}

func (OrderController OrderController) Update(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Update Called")

}

func (OrderController OrderController) Delete(app *gin.Context) {

	app.IndentedJSON(http.StatusOK, "Delete Called")

}

func Paginate(app *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page, _ := strconv.Atoi(app.DefaultQuery("page", "0"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(app.DefaultQuery("page_size", "10"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Filter(app *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		page, _ := strconv.Atoi(app.DefaultQuery("page", "0"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(app.DefaultQuery("page_size", "10"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
