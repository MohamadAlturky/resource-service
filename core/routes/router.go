package routes

import (
	"github.com/MohamadAlturky/Resources/core/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/set", services.SetHandler)
	router.GET("/get/:activityId", services.GetHandler)

	return router
}
