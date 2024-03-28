package product_sale_yearly_report

import (
	"math/big"
	"time"
)

type BigFloat struct {
	*big.Float
}

type ProductSaleYearlyReport struct {
	ID              int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	ProductID       uint      `json:"product_id"`
	CountSales      int       `json:"count_sales"`
	SellingPrice    float64   `json:"selling_price"`
	BuyingPrice     float64   `json:"buying_price"`
	TotalGrossSales float64   `json:"total_gross_sales"`
	TotalNettSales  float64   `json:"total_nett_sales"`
	Year            int       `json:"year"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
