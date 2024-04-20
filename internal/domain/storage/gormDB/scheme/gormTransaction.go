package scheme

type Transaction struct {
	Id            uint `gorm:"primaryKey;<-:create"`
	WarehausId    uint
	SoldPointId   uint
	MerchandiseId uint
	Date          string
	Count         uint
}
