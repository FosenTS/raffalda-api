package entity

type Transaction struct {
	Id            uint
	WarehouseId   uint
	SoldPointId   uint
	MerchandiseId uint
	Count         uint
}

type TransactionInfo struct {
	Id               uint
	WarehausId       uint
	SoldPointId      uint
	MerchandiseId    uint
	Count            uint
	WarehouseName    string
	SoldPointRegion  string
	SoldPointName    string
	SoldPointAddress string
	ProductName      string
	ProductCost      float64
	ManufactureDate  string
	ExpireDate       string
	SKU              int
	Quantity         uint
	Measure          string
}
