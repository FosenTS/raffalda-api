package entity

type MerchandiseMoreInfo struct {
	Id              uint
	WarehouseId     uint
	ProductName     string
	ProductCost     float64
	ManufactureDate string
	ExpiryDate      string
	SKU             int
	Quantity        uint
	Measure         string

	ExpireProcent uint
}
