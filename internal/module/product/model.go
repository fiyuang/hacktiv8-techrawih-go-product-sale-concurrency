package product

import (
	"math/big"
	"time"
)

type BigFloat struct {
	*big.Float
}
type Product struct {
	ID           int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name         string    `json:"name"`
	Stock        int       `json:"stock"`
	SellingPrice float64   `json:"selling_price"`
	BuyingPrice  float64   `json:"buying_price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
