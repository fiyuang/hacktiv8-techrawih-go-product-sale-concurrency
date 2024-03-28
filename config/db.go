package config

import (
	"fmt"
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/module/product"
	product_sale_yearly_report "hacktiv8-techrawih-go-product-sale-concurrency/internal/module/product-sale-yearly-report"
	"hacktiv8-techrawih-go-product-sale-concurrency/internal/module/sales"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBCon ... Database Connection Instance
var dbCon *gorm.DB

// InitDB ... function to initialize database
func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, dbport, user, dbname, password)

	dbCon, err = gorm.Open(postgres.Open(config), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}) // Adjust log level as needed
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	dbCon.Debug().AutoMigrate(product.Product{}, sales.Sales{}, product_sale_yearly_report.ProductSaleYearlyReport{})

}

func GetDBConnection() *gorm.DB {
	InitDB()

	return dbCon
}

func Migrate() *gorm.DB {
	InitDB()
	dbCon.Debug().AutoMigrate(product.Product{}, sales.Sales{}, product_sale_yearly_report.ProductSaleYearlyReport{})

	return dbCon
}
