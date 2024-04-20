package dto

type TransactionCreate struct {
	WarehousesId  uint
	SoldPointId   uint
	MerchandiseId uint
	Count         uint
}
