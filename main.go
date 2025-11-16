package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initDB connects to the database and runs auto-migrations.
func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate creates/updates tables based on structs (Message struct is visible from model.go)
	db.AutoMigrate(&Message{})
	return db
}

func main() {
	// 1. Initialize dependencies
	db := initDB()
	// Inject DB connection into our handler
	h := Handler{db: db}

	// 2. Create Echo server instance
	e := echo.New()

	// 3. Register routes
	// All handler methods are visible from handler.go
	e.GET("/messages", h.getHandler)
	e.POST("/messages", h.postHandler)
	e.PATCH("/messages/:id", h.patchHandler)
	e.DELETE("/messages/:id", h.deleteHandler)

	// 4. Start the server
	e.Start(":8080")
}
