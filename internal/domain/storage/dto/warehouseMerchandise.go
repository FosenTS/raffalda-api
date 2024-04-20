package dto

type WarehouseMerchandiseCreate struct {
	WarehouseId     uint
	ProductName     string
	ProductCost     float64
	ManufactureDate string
	ExpiryDate      string
	SKU             int
	Quantity        uint
	Measure         string
}
