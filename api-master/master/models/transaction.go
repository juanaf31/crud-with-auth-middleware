package models

type Transaction struct {
	OrderID      string `json:"order_id"`
	ProductName  string `json:"product_name"`
	CategoryName string `json:"category_name"`
	ProductCode  string `json:"product_code"`
	Price        string `json:"price"`
	Quantity     string `json:"quantity"`
	OrderDate    string `json:"order_date"`
	OutletName   string `json:"outlet_name"`
	RegionName   string `json:"region_name"`
	TotalPrice   int    `json:"total_price"`
}

type Outlet struct {
	OutletCode string `json:"outlet_code"`
	OutletName string `json:"outlet_name"`
	Region     string `json:"regional"`
}

type ProductPrice struct {
	ProductPriceID string `json:"product_price_id"`
	ProductPrice   string `json:"product_price"`
}
