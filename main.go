package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// RequestBody is a struct to bind JSON data
type RequestBody struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Database connection instance
var db *sql.DB

func main() {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Database connection
	var err error
	db, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Error opening the database:", err)
	} else {
		fmt.Println("Database connection successful")
	}

	// Check database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	} else {
		fmt.Println("Database connection successful")
	}

	defer db.Close()

	// Check and create table if not exists
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS items (
		id UUID PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		count INTEGER NOT NULL
	);
	`
	fmt.Println("Creating table...")
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	fmt.Println("Starting the server...")
	r := gin.Default()

	// /Register new item
	r.POST("/register", func(c *gin.Context) {
		var requestBody struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newID := uuid.New()
		_, err := db.Exec("INSERT INTO items (id, name, count) VALUES ($1, $2, $3)", newID, requestBody.Name, 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": newID.String()})
	})

	// /Add count to item
	r.POST("/add", func(c *gin.Context) {
		var requestBody struct {
			ID    string `json:"id"`
			Count int    `json:"count"`
		}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := uuid.Parse(requestBody.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		// Query data from the database
		var count int
		err = db.QueryRow("SELECT count FROM items WHERE id = $1", id).Scan(&count)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		// Update count
		newCount := count + requestBody.Count
		_, err = db.Exec("UPDATE items SET count = $1 WHERE id = $2", newCount, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Count updated"})
	})

	// Get data by ID
	r.GET("/item/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Query data from the database
		var item RequestBody
		err := db.QueryRow("SELECT name, count FROM items WHERE id = $1", id).Scan(&item.Name, &item.Count)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, item)
	})

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatal("Error starting the server:", err)
		}
	}()

	select {}
}
