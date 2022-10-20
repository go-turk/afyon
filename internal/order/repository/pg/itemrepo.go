package pg

import (
	"github.com/go-turk/afyon/internal/order/domain"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) CreateTable() error {
	return r.db.AutoMigrate(&domain.Item{})
}

func (r *ItemRepository) Save(item *domain.Item) error {
	return r.db.Create(&item).Error
}

func (r *ItemRepository) FindAll(ID ...int) ([]domain.Item, error) {
	var items []domain.Item
	return items, r.db.Find(&items, ID).Error
}

func (r *ItemRepository) FindByID(id int) (*domain.Item, error) {
	var item domain.Item
	return &item, r.db.First(&item, id).Error
}

func (r *ItemRepository) Delete(id int) error {
	return r.db.Delete(&domain.Item{}, id).Error
}

func (r *ItemRepository) Update(item domain.Item) error {
	return r.db.Save(&item).Error
}
