package repositories

import (
	"backend-speaker-clone/internal/models"
	"time"

	"gorm.io/gorm"
)

type ICustomersRepository interface {
	Create(item *models.CustomerInfo) error
	Update(item *models.CustomerInfo) error
	FindFirst(*models.CustomerInfo) error
	Count(item *models.CustomerInfo) (int64, error)
	FindList(item *models.CustomerInfo, page models.Page) ([]models.CustomerInfo, int64, error)
	Delete(item *models.CustomerInfo) error
}

type CustomersRepository struct {
	db *gorm.DB
}

func NewCustomersRepository(db *gorm.DB) *CustomersRepository {
	return &CustomersRepository{
		db: db,
	}
}

func (r *CustomersRepository) Create(item *models.CustomerInfo) error {
	now := time.Now()
	item.CreatedAt = now
	item.UpdatedAt = now
	return r.db.Create(item).Error
}

func (r *CustomersRepository) Update(item *models.CustomerInfo) error {
	item.UpdatedAt = time.Now()
	err := r.db.Save(item).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomersRepository) FindFirst(item *models.CustomerInfo) error {
	return r.db.Where(item).First(item).Error
}

func (r *CustomersRepository) Count(item *models.CustomerInfo) (int64, error) {
	db := r.db.Model(models.CustomerInfo{})
	total := int64(0)
	db = db.Where(item)
	db = db.Count(&total)
	return total, db.Error
}

func (r *CustomersRepository) FindList(item *models.CustomerInfo, page models.Page) ([]models.CustomerInfo, int64, error) {
	db := r.db.Model(models.CustomerInfo{})
	list := []models.CustomerInfo{}
	total := int64(0)
	db = db.Where(item)
	db.Count(&total)
	return list, total, db.Error
}
