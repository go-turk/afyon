package pg

import (
	"github.com/go-turk/afyon/internal/order/domain"
	"gorm.io/gorm"
	"log"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateTable() error {
	return r.db.AutoMigrate(&domain.Order{})
}

func (r *OrderRepository) Save(order *domain.Order) error {
	log.Println("order", order)
	return r.db.Create(&order).Error
}

func (r *OrderRepository) FindAll(ID ...int) ([]domain.Order, error) {
	var orders []domain.Order
	if len(ID) == 0 {
		return orders, r.db.Model(&orders).Preload("Items").Find(&orders).Error
	}
	return orders, r.db.Model(&orders).Preload("Items").Find(&orders, ID).Error
}

func (r *OrderRepository) FindByID(id int) (domain.Order, error) {
	var order domain.Order
	return order, r.db.Preload("Items").First(&order, id).Error
}

func (r *OrderRepository) AppendItem(orderID int, itemID int) error {
	var order domain.Order
	var item domain.Item
	return r.db.Model(&order).Association("Items").Append(&item)
}

func (r *OrderRepository) DecreaseItem(orderID int, itemID int) error {
	var order domain.Order
	var item domain.Item
	return r.db.Model(&order).Association("Items").Delete(&item)
}

func (r *OrderRepository) Delete(id int) error {
	return r.db.Delete(&domain.Order{}, id).Error
}

func (r *OrderRepository) Update(order domain.Order) error {
	return r.db.Save(&order).Error
}
