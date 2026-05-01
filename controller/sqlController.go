package controller

import (
	"ai-sql-customer-service/internal/llm"
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
	// 调用 AI 生成 SQL
	sqlText, err := llm.GenerateSQL(req.Question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 返回 AI 生成的 SQL
	c.JSON(http.StatusOK, model.AskResponse{
		Question: req.Question,
		SQL:      sqlText,
	})
}
