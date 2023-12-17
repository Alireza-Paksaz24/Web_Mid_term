package handler

import (
	"net/http"
	"strconv"
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

// Helper function to find a basket by ID
func findBasketByID(id int) (*Basket, int) {
	for i, basket := range b {
		if *basket.id == id {
			return &basket, i
		}
	}
	return nil, -1
}

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

// POST function for Handler to create a new basket
func (h *Handler) CreateBasket(c echo.Context) error {
	var request struct {
		Data  string `json:"data"`
		State bool   `json:"state"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	createdAt := time.Now()
	newBasket := createBasket(len(b)+1, createdAt, request.Data, request.State)

	return c.JSON(http.StatusCreated, newBasket)
}

// PATCH function for Handler to update a basket
func (h *Handler) UpdateBasket(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid basket ID"})
	}

	var request struct {
		Data  string `json:"data"`
		State bool   `json:"state"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	basketToUpdate, index := findBasketByID(id)
	if basketToUpdate == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "basket not found"})
	}

	basketToUpdate.data = request.Data
	basketToUpdate.state = &request.State
	basketToUpdate.update_at = time.Now()

	b[index] = *basketToUpdate

	return c.JSON(http.StatusOK, basketToUpdate)
}
