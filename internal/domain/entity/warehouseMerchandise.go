package entity

type WarehouseMerchandise struct {
	Id              uint
	WarehouseId     uint
	WarehouseName   string
	ProductName     string
	ProductCost     float64
	ManufactureDate string
	ExpireDate      string
	SKU             int
	Quantity        uint
	Measure         string
}
