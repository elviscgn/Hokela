package handlers

import (
	"time"
)

type Runner struct {
	// PRIMARY KEY
	RunnerID string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`

	// --- HOME BASE (Static) ---
	// This is where they live.
	// usage: "Find runners who live in Soweto."
	HomeLatitude  float64 `json:"home_latitude"`
	HomeLongitude float64 `json:"home_longitude"`

	// --- LIVE GPS (Dynamic) ---
	// This is where they are RIGHT NOW.
	// usage: "Find the closest runner to the shop."
	// Note: A string (e.g. "-26.204, 28.047") is harder to calculate distance with
	// than float64, but it works for simple storage.
	CurrentLocation string `json:"current_location"`

	// CONTACTS
	Email    string `json:"email" gorm:"unique;not null"`
	Contacts string `json:"contacts"`

	// HISTORY
	DeliveryHistory []Order `json:"delivery_history" gorm:"foreignKey:RunnerID"`

	// --- THE RATINGS EXPLAINED ---

	// 1. Rating FROM Customer (Reputation A)
	// "Did the runner smile? Was the food hot?"
	// The customer gives this score to the runner.
	RatingByCustomer float64 `json:"rating_by_customer" gorm:"default:5.0"`

	// 2. Rating FROM Spaza (Reputation B)
	// "Did the runner arrive on time to pick up? were they rude to staff?"
	// The Shop Owner gives this score to the runner.
	RatingBySpaza float64 `json:"rating_by_spaza" gorm:"default:5.0"`

	// VEHICLE
	VehicleType string `json:"vehicle_type"`

	// STATUS
	IsOnline bool `json:"is_online" gorm:"default:false"`

	// TIMESTAMPS
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
