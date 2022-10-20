package domain

import (
	"time"
)

type Item struct {
	ID        uint      `json:"ID" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	SKU       string    `json:"sku" gorm:"uuid,unique"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
}

func NewItem(name string, price float64) *Item {
	return &Item{
		Name:  name,
		Price: price,
	}
}
