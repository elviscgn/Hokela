package handlers

import (
	"time"
)

type Review struct {
	// PRIMARY KEY
	ReviewID string `json:"id" gorm:"primaryKey"`

	// THE AUTHOR (Who wrote this?)
	ReviewerID string `json:"reviewer_id" gorm:"index"`

	ReviewerName string `json:"reviewer_name"`

	Rating float64 `json:"rating"`
	Review string  `json:"review"`

	// --- THE TARGETS ---

	OrderID string `json:"order_id" gorm:"index"`

	// If this review is for the shop, this ID is filled.
	SpazaID string `json:"spaza_id" gorm:"index"`

	// 3. If this review is for the driver, this ID is filled.
	RunnerID string `json:"runner_id" gorm:"index"`

	// TIMESTAMPS
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
