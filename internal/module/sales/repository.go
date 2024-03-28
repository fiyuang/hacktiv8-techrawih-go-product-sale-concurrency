package sales

import (
	"context"
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/module/product"
	"log"
	"sync"

	"gorm.io/gorm"
)

// Repository Interface
type Repository interface {
	Save(ctx context.Context, request Sales) (*Sales, error)
	SaveAll(request []*Sales) error
	DeleteAll() error
	GetAll() ([]*Sales, error)
	GetProductByName(name string) (product.Product, error)
}

// NewRepository will implement SalesRepository Interface
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Save(ctx context.Context, request Sales) (res *Sales, err error) {
	dtoSale := Sales{ProductID: uint(request.ProductID), QtySold: request.QtySold, SaleAt: request.SaleAt}
	err = r.db.Save(&dtoSale).Error
	if err != nil {
		return nil, err
	}

	return &dtoSale, nil
}

func (r *repository) GetAll() ([]*Sales, error) {
	var salesDatas []*Sales

	if err := r.db.Find(&salesDatas).Error; err != nil {
		log.Fatalf("failed to get all records: %v", err)
		return nil, err
	}
	return salesDatas, nil
}

func (r *repository) DeleteAll() (err error) {
	if err = r.db.Where("1=1").Unscoped().Delete(&Sales{}).Error; err != nil {
		log.Fatalf("failed to delete records: %v", err)
		return err
	}
	return nil
}

func (r *repository) SaveAll(salesDto []*Sales) (err error) {
	var wg sync.WaitGroup
	batchSize := 500
	errorChannel := make(chan error, len(salesDto)/batchSize+1)

	for i := 0; i < len(salesDto); i += batchSize {
		end := i + batchSize
		if end > len(salesDto) {
			end = len(salesDto)
		}

		wg.Add(1)
		batch := make([]*Sales, end-i)
		copy(batch, salesDto[i:end])

		go func(b []*Sales) {
			defer wg.Done()
			// Memanggil saveBatch yang seharusnya menangani insert ke database
			if err := r.saveBatch(b); err != nil {
				errorChannel <- err
			}
		}(batch)
	}

	wg.Wait()
	close(errorChannel)

	for err = range errorChannel {
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) saveBatch(batch []*Sales) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Create(batch).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *repository) GetProductByName(name string) (res product.Product, err error) {
	query := r.db

	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}

	result := query.First(&res)

	return res, result.Error
}
