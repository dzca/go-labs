package repository

import (
	"example.com/domain"

	"gorm.io/gorm"
)

type driverRepo struct {
    db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) domain.DriverRepository {
    return &driverRepo{db}
}

func (r *driverRepo) Create(driver *domain.Driver) error { return r.db.Create(driver).Error }

func (r *driverRepo) GetByID(id uint) (*domain.Driver, error) {
    var driver domain.Driver
    err := r.db.First(&driver, id).Error
    return &driver, err
}
// func (r *driverRepo) Update(driver *domain.Driver) error { return r.db.Save(driver).Error }

func (r *driverRepo) Delete(id uint) error               { return r.db.Delete(&domain.Driver{}, id).Error }
func (r *driverRepo) List() ([]domain.Driver, error) {
    var drivers []domain.Driver
    err := r.db.Find(&drivers).Error
    return drivers, err
}

// Update driver by ID
func (r *driverRepo) Update(id uint, data map[string]interface{}) error {
    // This performs an UPDATE drivers SET ... WHERE id = id
    return r.db.Model(&domain.Driver{}).Where("id = ?", id).Updates(data).Error
}

