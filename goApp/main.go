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
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		count INTEGER NOT NULL,
		card_id VARCHAR(255) UNIQUE NOT NULL
	);	
	`

	fmt.Println("Creating table...")
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	fmt.Println("Starting the server...")
	r := gin.Default()

	// Register new item
	r.POST("/register", func(c *gin.Context) {
		var requestBody struct {
			Name   string `json:"name"`
			CardID string `json:"card_id"`
		}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var newID int
		err := db.QueryRow("INSERT INTO items (name, count, card_id) VALUES ($1, $2, $3) RETURNING id", requestBody.Name, 0, requestBody.CardID).Scan(&newID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": newID})
	})

	// Add count to item
	r.POST("/add", func(c *gin.Context) {
		var requestBody struct {
			ID     string `json:"id"`
			Count  int    `json:"count"`
			CardID string `json:"card_id"`
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
		var storedCardID string
		err = db.QueryRow("SELECT count, card_id FROM items WHERE id = $1", id).Scan(&count, &storedCardID)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		// Check if card_id matches
		if storedCardID != requestBody.CardID {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
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
	r.POST("/data", func(c *gin.Context) {
		var requestBody struct {
			CardID string `json:"card_id"` // Removed ID field
		}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Query data from the database using CardID
		var item struct {
			Name  string `json:"name"`
			Count int    `json:"count"`
		}
		err := db.QueryRow("SELECT name, count FROM items WHERE card_id = $1", requestBody.CardID).Scan(&item.Name, &item.Count)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		// Since we're using CardID directly, no need to check if it matches
		c.JSON(http.StatusOK, item)
	})

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatal("Error starting the server:", err)
		}
	}()

	select {}
}
