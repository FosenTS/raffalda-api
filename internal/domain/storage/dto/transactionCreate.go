package dto

type TransactionCreate struct {
	WarehousesId  uint
	SoldPointId   uint
	MerchandiseId uint
	Date          string
	Count         uint
}
