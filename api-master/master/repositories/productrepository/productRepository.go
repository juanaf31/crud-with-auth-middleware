package productrepository

import "liveCodeAPI/api-master/master/models"

type ProductRepository interface {
	GetAll() ([]*models.Product, error)
	GetByID(id string) (*models.Product, error)
	Delete(string) error
	Add([]*models.Product) ([]*models.Product, error)
	Update(string, *models.Product) (*models.Product, error)
}
