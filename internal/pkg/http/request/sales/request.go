package sales

import (
	"mime/multipart"
	"time"
)

type Import struct {
	File     multipart.FileHeader `form:"file" binding:"required"`
	FilePath string
}

type TempImportSales struct {
	ProductID   int
	ProductName string
	QtySold     int
	SaleAt      time.Time
}
