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

type WarehouseMerchandise struct {
	Id              uint
	WarehouseId     uint
	ProductName     string
	ProductCost     float64
	ManufactureDate string
	ExpireDate      string
	SKU             int
	Quantity        uint
	Measure         string
}
