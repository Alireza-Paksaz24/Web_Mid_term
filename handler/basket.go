package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Handler struct{}

type Basket struct {
	ID        *int      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Data      string    `json:"data"`
	State     *bool     `json:"state"`
}

var b = []Basket{}

// Helper function to string Array
func toString(array []Basket) string {
	str := "{"
	for _, i := range b {
		str += fmt.Sprintf("[%d,%s,%s,%s,%b]", *i.ID, i.CreatedAt, i.UpdatedAt, i.Data, *i.State)
	}
	str += "}"
	return str
}

// Helper function to find a basket by ID
func findBasketByID(id int) (*Basket, int) {
	for i, basket := range b {
		if *basket.ID == id {
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
	jsonData, _ := json.Marshal(b)
	return c.JSON(http.StatusOK, string(jsonData))
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

	return c.JSON(http.StatusCreated, newBasket.ID)
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

	basketToUpdate.Data = request.Data
	basketToUpdate.State = &request.State
	basketToUpdate.UpdatedAt = time.Now()

	b[index] = *basketToUpdate

	return c.JSON(http.StatusOK, basketToUpdate)
}

// GET function for Handler to retrieve a specific basket by ID
func (h *Handler) GetBasketByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid basket ID"})
	}

	basket, _ := findBasketByID(id)
	if basket == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "basket not found"})
	}

	return c.JSON(http.StatusOK, basket)
}

// DELETE function for Handler to delete a basket by ID
func (h *Handler) DeleteBasketByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid basket ID"})
	}

	basketToDelete, index := findBasketByID(id)
	if basketToDelete == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "basket not found"})
	}

	b = append(b[:index], b[index+1:]...)

	return c.JSON(http.StatusOK, map[string]string{"message": "basket deleted"})
}
