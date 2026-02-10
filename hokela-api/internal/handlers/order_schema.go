package handlers

import (
	"time"
)

type Order struct {
	// PRIMARY KEY
	OrderID string `json:"id" gorm:"primaryKey"`

	// FINANCIALS
	OrderPrice    float64 `json:"order_price"`
	DeliveryPrice float64 `json:"delivery_price"`

	// THE RECEIPT (Crucial Change)
	// This tells GORM: "This order has many receipt lines."
	Items []OrderItem `json:"items" gorm:"foreignKey:OrderID"`

	// LOGISTICS
	OrderDurationMinutes int `json:"order_duration_minutes"`

	// FOREIGN KEYS (The Links)
	SpazaID    string `json:"spaza_id" gorm:"index"`
	CustomerID string `json:"customer_id" gorm:"index"`
	RunnerID   string `json:"runner_id" gorm:"index"`

	// OPTIONAL: Link to a review
	//OrderReviewId string `json:"order_review_id" gorm:"index"`

	// TIMESTAMPS
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID uint `json:"id" gorm:"primaryKey"`

	// Link back to the Order
	OrderID string `json:"order_id" gorm:"index"`

	// Link to the original product (so you can click it to buy again)
	ProductID string `json:"product_id"`

	// SNAPSHOTS (The "Freeze")
	// We save the Name and Price here explicitly.
	// If the shop changes the price tomorrow, this order record stays correct.
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
