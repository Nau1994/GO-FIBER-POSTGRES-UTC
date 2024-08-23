package routes

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupRoutesApp() *fiber.App {
	app := fiber.New()
	UserRoutes(app)
	return app
}

func TestUserRoutes(t *testing.T) {
	app := setupRoutesApp()

	// Test POST /users
	user := `{"name":"Test User","email":"test.user@example.com","balance":100}`
	req := httptest.NewRequest("POST", "/users", bytes.NewReader([]byte(user)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	assert.Equal(t, 201, resp.StatusCode)

	// Test GET /users
	req = httptest.NewRequest("GET", "/users", nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}
