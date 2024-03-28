package sales

import (
	"time"
)

type Sales struct {
	ID        uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ProductID uint       `json:"product_id"`
	QtySold   int        `json:"qty_sold"`
	SaleAt    *time.Time `json:"sale_at"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
