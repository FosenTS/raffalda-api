package entity

type MerchandiseMoreInfo struct {
	Id              uint
	WarehouseId     uint
	WarehouseName   string
	ProductName     string
	ProductCost     float64
	ManufactureDate string
	ExpiryDate      string
	SKU             int
	Quantity        uint
	Measure         string

	ExpirePercentage uint
}
