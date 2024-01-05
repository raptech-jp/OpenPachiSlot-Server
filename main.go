package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// RequestBody is a struct to bind JSON data
type RequestBody struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// データベース接続
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	} else {
		fmt.Println("Success connecting to the database:", db)
	}
	defer db.Close()

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

	r.POST("/add", func(c *gin.Context) {
		// JSONデータを構造体にバインド
		var requestBody RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// バインドされたデータを使用
		message := "Hello " + requestBody.Name + ", count: " + fmt.Sprint(requestBody.Count)
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	// 別のgoroutine内でGinを実行
	go func() {
		if err := r.Run(":8080"); err != nil {
			fmt.Println("Error starting the server:", err)
		}
	}()
	select {}
}
