package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Handler struct{}

type Basket struct {
	id        *int
	create_at time.Time
	update_at time.Time
	data      string
	state     *bool
}

var b = []Basket{}

// Function to create a Basket
func createBasket(id int, createdAt time.Time, data string, state bool) Basket {
	updatedAt := createdAt

	// Create a new instance of the Basket struct
	basket := Basket{
		&id,
		createdAt,
		updatedAt,
		data,
		&state,
	}
	b = append(b, basket)
	// Return the struct
	return basket
}

// Handler for the /baskets endpoint
func (h *Handler) GetBaskets(c echo.Context) error {
	return c.JSON(http.StatusOK, b)
}
