package main

import (
	"errors"
	"gorm.io/gorm"
)

type LockRepository interface {
	FindById(uint64) (*Product, error)
	UpdateWithVersion(*Product) error
}

type LockRepositoryImpl struct {
	db *gorm.DB
}

func NewLockRepositoryImpl(db *gorm.DB) LockRepository {
	return LockRepositoryImpl{
		db: db,
	}
}

func (repo LockRepositoryImpl) FindById(id uint64) (*Product, error) {
	var product Product
	if err := repo.db.Model(&Product{}).Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo LockRepositoryImpl) UpdateWithVersion(product *Product) error {
	result := repo.db.Model(&product).Where("version = ?", product.Version).Updates(
		&Product{
			SaleCount: product.SaleCount + 1,
			Version:   product.Version + 1,
		})

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected == 0 {
		return errors.New("version mismatched")
	}
	return nil
}
