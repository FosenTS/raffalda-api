package dto

type TransactionCreate struct {
	WarehouseId   uint
	SoldPointId   uint
	MerchandiseId uint
	Date          string
	Count         uint
}
