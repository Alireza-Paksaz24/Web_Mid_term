package main

import (
	"log"

	"github.com/Alireza-Paksaz24/Web_Mid_term/handler"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=mid_term port=5432 sslmode=disable TimeZone=Asia/Tehran"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// AutoMigrate will create the table if it doesn't exist
	db.AutoMigrate(&handler.Basket{})
}

func main() {
	e := echo.New()

	h := handler.NewHandler(db)

	// Register the GET function for retrieving all baskets
	e.GET("/basket", h.GetBaskets)

	// Register the POST function for creating a new basket
	e.POST("/basket", h.CreateBasket)

	// Register the PATCH function for updating a basket
	e.PATCH("/basket/:id", h.UpdateBasket)

	// Register the GET function for retrieving a specific basket by ID
	e.GET("/basket/:id", h.GetBasketByID)

	// Register the DELETE function for deleting a basket by ID
	e.DELETE("/basket/:id", h.DeleteBasketByID)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
