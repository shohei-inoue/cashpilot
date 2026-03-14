package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
			"status":  "ok",
		})
	})
	router.Run(":8080") // 8080番ポートでサーバーを起動
}
