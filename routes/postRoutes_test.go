package routes

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupPostRoutesApp() *fiber.App {
	app := fiber.New()
	PostRoutes(app)
	return app
}

func TestPostRoutes(t *testing.T) {
	app := setupPostRoutesApp()

	// Test POST /posts
	post := `{"title":"Test Post","content":"Post content","user_id":1}`
	req := httptest.NewRequest("POST", "/posts", bytes.NewReader([]byte(post)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	assert.Equal(t, 201, resp.StatusCode)

	// Test GET /posts
	req = httptest.NewRequest("GET", "/posts", nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	assert.Equal(t, 200, resp.StatusCode)
}
