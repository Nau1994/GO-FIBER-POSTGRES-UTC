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

func setupPostApp() *fiber.App {
	app := fiber.New()
	app.Post("/posts", CreatePost)
	app.Get("/posts", GetAllPosts)
	return app
}

func TestCreatePost(t *testing.T) {
	app := setupPostApp()

	// Test creating a post
	post := models.Post{Title: "New Post", Content: "Post content1", UserID: 1}
	body, _ := json.Marshal(post)
	req := httptest.NewRequest("POST", "/posts", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, 201, resp.StatusCode)

	// Verify response
	var respPost models.Post
	if err := json.NewDecoder(resp.Body).Decode(&respPost); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	assert.NotZero(t, respPost.ID)
	assert.Equal(t, "New Post", respPost.Title)
}

func TestGetAllPosts(t *testing.T) {
	app := setupPostApp()

	// Test getting all posts
	req := httptest.NewRequest("GET", "/posts", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)

	// Verify response
	var posts []models.Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	assert.NotEmpty(t, posts)
}
