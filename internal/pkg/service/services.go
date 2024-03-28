package service

import (
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/module/product"

	"gorm.io/gorm"
)

type Services struct {
	ProductService product.Repository
}

func Init(db *gorm.DB) *Services {
	return &Services{
		ProductService: product.NewRepository(db),
	}
}
