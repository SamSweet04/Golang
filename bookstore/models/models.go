package models

import "time"

// Book represents a book in the bookstore
type Book struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Cost        float32   `gorm:"not null" json:"cost"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}
