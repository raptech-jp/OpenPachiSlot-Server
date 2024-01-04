package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "world",
		})
	})
	r.GET("/call", func(c *gin.Context) {
		// クエリパラメータを取得
		name := c.Query("name")
		c.JSON(200, gin.H{
			"message": "Hello " + name,
		})
	})
	r.Run()
}
