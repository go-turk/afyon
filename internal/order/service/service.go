package service

import (
	"github.com/go-turk/afyon/internal/order"
	"github.com/go-turk/afyon/internal/order/domain"
)

type OrderService struct {
	itemRepo  order.ItemRepository
	orderRepo order.OrderRepository
}

func NewOrderService(itemRepository order.ItemRepository, orderRepository order.OrderRepository) *OrderService {
	itemRepository.CreateTable()
	orderRepository.CreateTable()

	return &OrderService{
		itemRepo:  itemRepository,
		orderRepo: orderRepository,
	}
}

func (s *OrderService) Create(order *domain.Order) error {
	return s.orderRepo.Save(order)
}

func (s *OrderService) FindAll(ID ...int) ([]domain.Order, error) {
	return s.orderRepo.FindAll(ID...)
}

func (s *OrderService) FindByID(id int) (domain.Order, error) {
	return s.orderRepo.FindByID(id)
}

func (s *OrderService) AppendItem(orderID int, itemID int) error {
	return s.orderRepo.AppendItem(orderID, itemID)
}

func (s *OrderService) DecreaseItem(orderID int, itemID int) error {
	return s.orderRepo.DecreaseItem(orderID, itemID)
}

func (s *OrderService) CreateItem(item *domain.Item) error {
	return s.itemRepo.Save(item)
}

func (s *OrderService) GetItem(id int) (*domain.Item, error) {
	return s.itemRepo.FindByID(id)
}

func (s *OrderService) FindItems(ID ...int) ([]domain.Item, error) {
	return s.itemRepo.FindAll(ID...)
}
