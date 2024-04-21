package entity

type Warehouse struct {
	ID       uint
	Name     string
	Volume   uint
	Capacity uint
}

type WarehouseMapInfo struct {
	Monday    uint
	Tuesday   uint
	Wednesday uint
	Thursday  uint
	Friday    uint
	Saturday  uint
	Sunday    uint

	Top1 *MerchandiseMoreInfo
	Top2 *MerchandiseMoreInfo
	Top3 *MerchandiseMoreInfo
	Top4 *MerchandiseMoreInfo

	WarehouseName string
	Volume        uint
	Capacity      uint
}
