package categoryrepository

import "liveCodeAPI/api-master/master/models"

type CategoryRepository interface {
	GetAll() ([]*models.Category, error)
	GetByID(string) (*models.Category, error)
	Delete(string) error
	Add([]*models.Category) ([]*models.Category, error)
	Update(string, *models.Category) (*models.Category, error)
}
