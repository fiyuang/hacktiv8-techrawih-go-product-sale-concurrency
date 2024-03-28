package product

import (
	"context"
	"gorm.io/gorm"
	"log"
)

// Repository Interface
type Repository interface {
	Insert(ctx context.Context, product Product) (Product, error)
	GetIdByName(name string) (Product, error)
	GetAllProduct() ([]*Product, error)
}

// NewRepository will implement ProductRepository Interface
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Insert(ctx context.Context, product Product) (Product, error) {
	result := r.db.Save(&product)
	return product, result.Error
}

func (r *repository) GetIdByName(name string) (Product, error) {
	var res Product
	query := r.db

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	result := query.First(&res)

	return res, result.Error
}

func (r *repository) GetAllProduct() ([]*Product, error) {
	var productDatas []*Product

	if err := r.db.Find(&productDatas).Error; err != nil {
		log.Fatalf("failed to get all records: %v", err)
		return nil, err
	}
	return productDatas, nil
}
