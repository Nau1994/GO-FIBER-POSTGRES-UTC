package controllers

import (
	"GoFiber/config"
	"GoFiber/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	fmt.Println(config.DB)
	if err := config.DB.Create(&post).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(post)
}

func GetAllPosts(c *fiber.Ctx) error {
	var posts []models.Post
	if err := config.DB.Find(&posts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(posts)
}
