package handlers

import (
	"time"
)

type Product struct {
	// PRIMARY KEY
	ProductID string `json:"id" gorm:"primaryKey"`

	// INFO
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`

	// THE PARENT (The Shop)
	SpazaID string `json:"spaza_id" gorm:"index"`

	// TIMESTAMPS
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
