package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `gorm:"foreignKey:UserID"`
}
