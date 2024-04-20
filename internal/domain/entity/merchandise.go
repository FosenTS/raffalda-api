package entity

type Merchandise struct {
	ID              uint
	WarehouseID     uint
	Warehouse       Warehouse
	ProductName     string
	ProductCost     float64
	ManufactureDate string
	ExpiryDate      string
	SKU             int
	StoreName       string
	StoreAddress    string
	Region          string
	SaleDate        string
	QuantitySold    float64
	ProductAmount   float64
	ProductMeasure  float64
	Manufacturer    string
}
