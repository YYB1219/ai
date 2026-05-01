package main

import (
	"ai-sql-customer-service/router"
)

func main() {
	r := router.SetupRouter()
	// 启动 HTTP 服务
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
