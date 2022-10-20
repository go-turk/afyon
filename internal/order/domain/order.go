package domain

import (
	"time"
)

const (
	// StatusCreated ...
	StatusCreated = "created"
	// StatusPaid ...
	StatusPaid = "paid"
	// StatusShipped ...
	StatusShipped = "shipped"
	// StatusDelivered ...
	StatusDelivered = "delivered"
)

type Order struct {
	ID        uint      `json:"ID" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	Items     []Item    `json:"items" gorm:"many2many:order_items"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
}

// NewOrder Method to create a new order
func NewOrder(items []Item, userID int) *Order {
	return &Order{
		Items:  items,
		UserID: userID,
		Status: StatusCreated,
	}
}
