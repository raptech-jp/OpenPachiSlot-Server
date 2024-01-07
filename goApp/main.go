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
	initializeDatabase()
	router := setupRouter()

	log.Println("Starting the server...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}

func initializeDatabase() {
	var err error
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	db, err = sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	createTable()
}

func createTable() {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS items (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        count INTEGER NOT NULL,
        card_id UUID UNIQUE NOT NULL
    );	
    `
	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())

	r.POST("/user", registerUser)
	r.PUT("/user/:uuid", updateUser)
	r.GET("/user/:uuid", getUserData)
	r.GET("/all-items", getAllItems)
	return r
}

// Register new item
func registerUser(c *gin.Context) {
	var requestBody struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	CardID := uuid.New().String()

	// Check if cardID already exists
	var existingID int
	err := db.QueryRow("SELECT id FROM items WHERE card_id = $1", CardID).Scan(&existingID)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registration failed. Please try again."})
		return
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var newID int
	err = db.QueryRow("INSERT INTO items (name, count, card_id) VALUES ($1, $2, $3) RETURNING id", requestBody.Name, 0, CardID).Scan(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": newID, "name": requestBody.Name, "cardId": CardID})
}

func updateUser(c *gin.Context) {
	var requestBody struct {
		ID  int `json:"id"`
		AddCount int `json:"add"`
	}
	CardID := c.Param("uuid")

	if _, err := uuid.Parse(CardID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := requestBody.ID

	// Query data from the database
	var storedCount int
	var storedCardID string
	var err error
	err = db.QueryRow("SELECT count, card_id FROM items WHERE id = $1", id).Scan(&storedCount, &storedCardID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Check if card_id matches
	if storedCardID != CardID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	// Update count
	newCount := storedCount + requestBody.AddCount
	_, err = db.Exec("UPDATE items SET count = $1 WHERE id = $2", newCount, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Count updated", "newCount": newCount})
}

// Get data by ID
func getUserData(c *gin.Context) {
	CardID := c.Param("uuid")

	// Query data from the database using CardID
	var item struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Count int    `json:"count"`
	}

	if _, err := uuid.Parse(CardID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	err := db.QueryRow("SELECT id, name, count FROM items WHERE card_id = $1", CardID).Scan(&item.ID, &item.Name, &item.Count)
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
}

func getAllItems(c *gin.Context) {
	var items []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Count  int    `json:"count"`
		CardID string `json:"card_id"`
	}

	rows, err := db.Query("SELECT id, name, count, card_id FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Count  int    `json:"count"`
			CardID string `json:"card_id"`
		}
		if err := rows.Scan(&item.ID, &item.Name, &item.Count, &item.CardID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// CORS middleware
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
