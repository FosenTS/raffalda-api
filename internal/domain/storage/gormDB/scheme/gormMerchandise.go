package scheme

import "gorm.io/gorm"

type Merchandise struct {
	gorm.Model
	ID              uint `gorm:"primaryKey;<-:create"`
	WarehouseID     uint
	Warehouse       Warehouse `gorm:"foreignKey:WarehouseID"`
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
