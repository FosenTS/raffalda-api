package entity

type WarehouseMoreInfo struct {
	Id           uint
	Name         string
	Volume       uint
	Capacity     uint
	Merchandises []*WarehouseMerchandise
}
