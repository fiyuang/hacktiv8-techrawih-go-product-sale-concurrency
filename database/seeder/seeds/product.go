package seeds

import (
	"context"
	"encoding/json"
	"fmt"
	"hacktiv8-techrawih-go-product-sale-concurrency/config"
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/module/product"
	"os"
	"time"
)

func (s Seed) Product(context context.Context, repository product.Repository) {
	var jsonFile []byte
	var err error
	db := config.GetDBConnection()

	jsonFile, err = os.ReadFile("database/seeder/seeds/json/product.json")
	if err != nil {
		fmt.Println("error when parse json file", err)
	}

	var products []product.Product
	err = json.Unmarshal(jsonFile, &products)
	if err != nil {
		fmt.Println("error while parse json file")
		return
	}

	for _, eachProduct := range products {
		productDto := product.Product{
			Name:         eachProduct.Name,
			Stock:        eachProduct.Stock,
			SellingPrice: eachProduct.SellingPrice,
			BuyingPrice:  eachProduct.BuyingPrice,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		err = db.Create(&productDto).Error
		if err != nil {
			fmt.Println("error while create seed data")
			return
		}
	}

}
