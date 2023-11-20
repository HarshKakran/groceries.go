package routers

import (
	"github.com/HarshKakran/groceries.go/handlers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	// Common routes
	router.POST("/api/signin", handlers.Signin)
	router.POST("/api/signup", handlers.Signup)
	router.GET("/api/signout", handlers.Signout)

	// Owner routes
	router.GET("/api/store", handlers.StoreInfo)
	router.POST("/api/store/item", handlers.CreateItem)
	router.PUT("/api/store/item/:id", handlers.UpdateItem)
	router.DELETE("/api/store/item/:id", handlers.DeleteItem)

	// Consumer routes
	router.GET("/api/consumer", handlers.ConsumerPage)
	router.POST("/api/consumer/checkout", handlers.CreateOrder)

	return router
}
