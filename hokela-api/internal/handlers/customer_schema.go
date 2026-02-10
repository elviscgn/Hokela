package handlers

import (
	"time"
)

type Customer struct {
	// PRIMARY KEY
	CustomerID string `json:"id" gorm:"primaryKey"`

	// BASIC INFO
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	// CONTACTS
	Email    string `json:"email" gorm:"unique;not null"`
	Contacts string `json:"contacts"`

	// PREFERENCES
	// Added default 'sms' so it doesn't break if left empty
	NotifPreference string `json:"notif_preference" gorm:"default:'sms'"`

	// HISTORY
	// This looks at the 'Order' table for rows that have 'customer_id' = this ID
	OrderHistory []Order `json:"order_history" gorm:"foreignKey:CustomerID"`

	// RATING
	CustomerRating float64 `json:"customer_rating" gorm:"default:5.0"`

	// TIMESTAMPS
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
