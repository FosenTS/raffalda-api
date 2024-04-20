package dto

import "raffalda-api/internal/domain/storage/gormDB/scheme"

type MerchandiseCreate struct {
	WarehouseID     uint
	Warehouse       scheme.Warehouse `gorm:"foreignKey:CompanyRefer"`
	ProductName     string
	ProductCost     float64
	ManufactureDate string
	ExpiryDate      string
	SKU             int
	StoreName       string
	StoreAddress    string
	Region          string
	SaleDate        string
	QuantitySold    float64
	ProductAmount   float64
	ProductMeasure  string
	ProductVolume   float64
	Manufacturer    string
}
