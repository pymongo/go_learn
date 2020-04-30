package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 在ab -n 200 -c 200的测试环境下，每秒能处理18000个请求，而actix只能处理13000个
	// 在ab -n 1000 -c 200的测试环境下，每秒能处理17000个请求，而actix也能处理17000个
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8102")
}
