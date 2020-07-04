package productusecase

import "liveCodeAPI/api-master/master/models"

type ProductUsecase interface {
	GetProducts() ([]*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
	DeleteProduct(id string) error
	AddProduct([]*models.Product) ([]*models.Product, error)
	UpdateProduct(string, *models.Product) (*models.Product, error)
}
