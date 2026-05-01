package controller

import (
	"ai-sql-customer-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func AIask(c *gin.Context) {
	var req model.AskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "请求参数错误",
		})
		return
	}

	if req.Question == "" { //提问是否为空
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "question 不能为空",
		})
		return
	}

	c.JSON(http.StatusOK, model.AskResponse{
		Question: req.Question,
		Message:  "已收到你的问题，下一步会生成 SQL",
	})
}
