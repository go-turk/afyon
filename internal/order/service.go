package order

import "github.com/go-turk/afyon/internal/order/domain"

type OrderService interface {
	AppendItem(orderID int, itemID int) error
	Create(order *domain.Order) error
	CreateItem(item *domain.Item) error
	DecreaseItem(orderID int, itemID int) error
	FindAll(ID ...int) ([]domain.Order, error)
	FindByID(id int) (domain.Order, error)
	FindItems(ID ...int) ([]domain.Item, error)
	GetItem(id int) (*domain.Item, error)
}
