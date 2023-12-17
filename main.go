package main

import (
	"log"

	"github.com/Alireza-Paksaz24/Web_Mid_term/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	h := handler.Handler{}

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
