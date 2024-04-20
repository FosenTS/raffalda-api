package entity

type Transaction struct {
	Id            uint
	WarehouseId   uint
	SoldPointId   uint
	MerchandiseId uint
	Date          string
	Count         uint
}

type TransactionInfo struct {
	Id               uint
	WarehausId       uint
	SoldPointId      uint
	MerchandiseId    uint
	Date             string
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

type TransactionStats struct {
	MaxValue  uint
	Monday    uint
	Tuesday   uint
	Wednesday uint
	Thursday  uint
	Friday    uint
	Saturday  uint
	Sunday    uint
}
