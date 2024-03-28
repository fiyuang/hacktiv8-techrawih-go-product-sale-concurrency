package sales

import (
	"errors"
	"fmt"
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/module/product"
	product_sale_yearly_report "hacktiv8-techrawih-go-product-sale-concurrency/internal/module/product-sale-yearly-report"
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/pkg/http/request/sales"
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/pkg/utils"
	"log"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Service interface {
	SaveAll(request sales.Import) (*product_sale_yearly_report.ProductSaleYearlyReport, error)
}

type service struct {
	repo                  Repository
	productRepo           product.Repository
	ProductSaleYearlyRepo product_sale_yearly_report.Repository
	db                    *gorm.DB
}

func NewService(repo Repository, productRepo product.Repository, productTrxRepo product_sale_yearly_report.Repository, db *gorm.DB) Service {
	return &service{
		repo:                  repo,
		productRepo:           productRepo,
		ProductSaleYearlyRepo: productTrxRepo,
		db:                    db,
	}
}

func (us *service) SaveAll(request sales.Import) (res *product_sale_yearly_report.ProductSaleYearlyReport, err error) {
	records, err := utils.ReadCSV(request.FilePath)
	if err != nil {
		return nil, err
	}

	if err = us.repo.DeleteAll(); err != nil {
		log.Fatalf("failed to delete all sales records: %v", err)
		return nil, err
	}

	if err = us.ProductSaleYearlyRepo.DeleteAll(); err != nil {
		log.Fatalf("failed to delete all product_trxes records: %v", err)
		return nil, err
	}

	productDatas, err := us.productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}

	var salesStoreDto []*Sales
	notExistProduct := []string{}

	for index, record := range records {
		if index != 0 {
			qtySold, _ := strconv.Atoi(record[1])
			saleAt, _ := utils.ConvertStringToTime(record[2])
			salesDto := &Sales{
				QtySold: qtySold,
				SaleAt:  saleAt,
			}

			productExists := false
			for _, productData := range productDatas {
				if productData.Name == record[0] {
					salesDto.ProductID = uint(productData.ID)
					productExists = true
					break
				}
			}
			if !productExists {
				fmt.Println("Product Not Found === ", record[0])
				notExistProduct = append(notExistProduct, record[0]+" - "+record[2])
			}

			salesStoreDto = append(salesStoreDto, salesDto)
		}
	}

	if len(notExistProduct) > 0 {
		errorMessage := "Import failed! Product(s) not found: " + strings.Join(notExistProduct, ", ")
		return nil, errors.New(errorMessage)
	}

	err = us.repo.SaveAll(salesStoreDto)
	if err != nil {
		return nil, err
	}

	ch := make(chan []*product_sale_yearly_report.ProductSaleYearlyReport)

	errCh := make(chan error)

	go product_sale_yearly_report.AggregateSalesByProduct(us.db, ch, errCh)
	select {
	case yearlyReports := <-ch:
		product_sale_yearly_report.UpdateStockProduct(us.db, yearlyReports, errCh)
	case err := <-errCh:
		// fmt.Println("Error aggregating sales:", err)
		return nil, err
	}

	return res, nil
}
