package product_sale_yearly_report

import (
	"time"

	"gorm.io/gorm"
)

func AggregateSalesByProduct(db *gorm.DB, ch chan []*ProductSaleYearlyReport, errCh chan error) {
	var results []struct {
		ProductID    uint    `gorm:"column:product_id"`
		QuantitySold float64 `gorm:"column:total_quantity"`
		SellingPrice float64 `gorm:"column:selling_price"`
		BuyingPrice  float64 `gorm:"column:buying_price"`
		SaleAt       int     `gorm:"column:sale_year"`
	}

	if err := db.Table("sales").
		Select("sales.product_id, EXTRACT(YEAR FROM sales.sale_at) AS sale_year, SUM(sales.qty_sold) as total_quantity, products.selling_price, products.buying_price").
		Joins("left join products on products.id = sales.product_id").
		Group("sales.product_id, EXTRACT(YEAR FROM sales.sale_at), products.selling_price, products.buying_price").
		Scan(&results).Error; err != nil {
		errCh <- err
		return
	}

	var yearlyReportDto []*ProductSaleYearlyReport
	for _, result := range results {
		totalGrossSales := result.QuantitySold * result.SellingPrice
		totalNettSales := totalGrossSales - (result.QuantitySold * result.BuyingPrice)

		productSaleYearlyReport := &ProductSaleYearlyReport{
			ProductID:       result.ProductID,
			SellingPrice:    result.SellingPrice,
			BuyingPrice:     result.BuyingPrice,
			TotalGrossSales: totalGrossSales,
			TotalNettSales:  totalNettSales,
			CountSales:      int(result.QuantitySold),
			Year:            result.SaleAt,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}

		yearlyReportDto = append(yearlyReportDto, productSaleYearlyReport)
	}

	// Use transaction for bulk inserts
	tx := db.Begin()
	if err := tx.Error; err != nil {
		errCh <- err
	}

	if err := tx.CreateInBatches(yearlyReportDto, len(yearlyReportDto)).Error; err != nil {
		tx.Rollback() // Rollback in case of any error
		errCh <- err
	}

	err := tx.Commit().Error
	if err != nil {
		errCh <- err
	}

	ch <- yearlyReportDto
}

func UpdateStockProduct(db *gorm.DB, yearlyReports []*ProductSaleYearlyReport, errCh chan error) {
	tx := db.Begin()
	if tx.Error != nil {
		errCh <- tx.Error
	}

	query := `
	UPDATE products
	SET stock = stock - (SELECT count_sales FROM product_sale_yearly_reports WHERE product_sale_yearly_reports.product_id = products.id)
	WHERE EXISTS (
		SELECT 1 FROM product_sale_yearly_reports WHERE product_sale_yearly_reports.product_id = products.id
	);
	`
	if err := tx.Exec(query).Error; err != nil {
		tx.Rollback()
		errCh <- err
	}

	if err := tx.Commit().Error; err != nil {
		errCh <- err
	}

	close(errCh)
}
