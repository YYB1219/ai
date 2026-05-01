package router

import (
	"ai-sql-customer-service/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controller.Test)
	r.POST("/ask", controller.AIask)
	return r
}
