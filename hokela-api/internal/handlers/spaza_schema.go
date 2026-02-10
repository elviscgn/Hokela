package handlers

import (
	"time"
)

type Spaza struct {
	// PRIMARY KEY
	SpazaID string `json:"id" gorm:"primaryKey"`

	// BASIC INFO
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Email     string  `json:"email" gorm:"unique;not null"`
	Picture   string  `json:"picture"`

	// BUSINESS DETAILS
	OperatingHours string `json:"operating_hours"`
	OwnerContacts  string `json:"owner_contacts"`

	// --- PAYMENT ---

	PayFastMerchantID string `json:"payfast_merchant_id"`

	// RELATIONSHIPS
	Inventory    []Product `json:"inventory" gorm:"foreignKey:SpazaID"`
	OrderHistory []Order   `json:"order_history" gorm:"foreignKey:SpazaID"`

	// TIMESTAMPS
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
