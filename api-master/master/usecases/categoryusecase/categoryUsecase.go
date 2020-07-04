package categoryusecase

import "liveCodeAPI/api-master/master/models"

type CategoryUsecase interface {
	GetCategories() ([]*models.Category, error)
	GetCategoryByID(id string) (*models.Category, error)
	DeleteCategory(id string) error
	AddCategory([]*models.Category) ([]*models.Category, error)
	UpdateCategory(string, *models.Category) (*models.Category, error)
}
