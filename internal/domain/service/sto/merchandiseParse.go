package sto

type MerchandiseParse struct {
	ProductName     string  `csv:"Product_Name"`
	ProductCost     float64 `csv:"Product_Cost"`
	ManufactureDate string  `csv:"Manufacture_Date"`
	ExpiryDate      string  `csv:"Expiry_Date"`
	SKU             int     `csv:"SKU"`
	StoreName       string  `csv:"Store_Name"`
	StoreAddress    string  `csv:"Store_Address"`
	Region          string  `csv:"Region"`
	SaleDate        string  `csv:"Sale_Date"`
	QuantitySold    float64 `csv:"Quantity_Sold"`
	ProductAmount   float64 `csv:"Product_Amount"`
	ProductMeasure  string  `csv:"Product_Measure"`
	ProductVolume   float64 `csv:"Product_Volume"`
	Manufacturer    string  `csv:"Manufacturer"`
}
