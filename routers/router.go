package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var orderRouter OrderRouter

func Init() *gin.Engine {
	// v1 route group
	router := gin.Default()

	router.Use(cors.Default())

	v1 := router.Group("/v1")
	{
		orderRouter.handle(v1.Group("/orders"))
	}

	return router
}
