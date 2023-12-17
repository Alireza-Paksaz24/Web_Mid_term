package main

import (
	"log"

	handler "github.com/Alireza-Paksaz24/Web_Mid_term"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	h := handler.Handler{}

	e.GET("/basket", h.GET)

	e.POST("/basket", h.CreateBasket)

	e.PATCH("/basket", h.UpdateBasket)

	e.GET("/basket", h.GetBasketByID)

	e.DELETE("/basket", h.DELETE)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
