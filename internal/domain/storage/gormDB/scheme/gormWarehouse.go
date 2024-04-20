package scheme

type Warehouse struct {
	ID       uint `gorm:"primaryKey;<-:create"`
	Name     string
	Volume   uint
	Capacity uint
}

type WarehouseMerchandise struct {
	Id              uint `gorm:"primaryKey;<-:create"`
	WarehouseId     uint
	ProductName     string
	ProductCost     float64
	ManufactureDate string
	ExpiryDate      string
	SKU             int
	Quantity        uint
	Measure         string
}
