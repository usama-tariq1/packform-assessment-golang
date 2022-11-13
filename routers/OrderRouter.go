package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/usama-tariq1/leet-gin/controllers"
)

type OrderRouter struct {
}

func (OrderRouter OrderRouter) handle(router *gin.RouterGroup) {
	var orderController controllers.OrderController
	router.GET("/", orderController.Index)
}
