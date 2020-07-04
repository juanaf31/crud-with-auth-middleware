package models

type Product struct {
	ID              string   `json:"product_id"`
	ProductCode     string   `json:"product_code"`
	ProductName     string   `json:"product_name"`
	ProductCategory Category `json:"product_category"`
}
