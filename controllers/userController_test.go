package controllers

import (
	"GoFiber/models"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	app := fiber.New()
	app.Post("/users", CreateUser)
	app.Get("/users", GetAllUsers)
	return app
}

func TestCreateUser(t *testing.T) {
	app := setupApp()

	// Test creating a user
	user := models.User{Name: "Jane Doe1", Email: "jane.doe@example.com", Balance: 100}
	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, 201, resp.StatusCode)

	// Verify response
	var respUser models.User
	if err := json.NewDecoder(resp.Body).Decode(&respUser); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	assert.NotZero(t, respUser.ID)
	assert.Equal(t, "Jane Doe1", respUser.Name)
}

func TestGetAllUsers(t *testing.T) {
	app := setupApp()

	// Test getting all users
	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)

	// Verify response
	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	assert.NotEmpty(t, users)
}
