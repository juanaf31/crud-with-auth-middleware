package categoryusecase

import (
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/repositories/categoryrepository"
)

type CategoryUsecaseImpl struct {
	categoryRepo categoryrepository.CategoryRepository
}

func InitCategoryUsecase(categoryRepo categoryrepository.CategoryRepository) CategoryUsecase {
	return &CategoryUsecaseImpl{categoryRepo: categoryRepo}
}

func (c *CategoryUsecaseImpl) GetCategories() ([]*models.Category, error) {
	categories, err := c.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryUsecaseImpl) GetCategoryByID(id string) (*models.Category, error) {
	category, err := c.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryUsecaseImpl) DeleteCategory(id string) error {
	err := c.categoryRepo.Delete(id)
	return err
}

func (c *CategoryUsecaseImpl) AddCategory(data []*models.Category) ([]*models.Category, error) {
	category, err := c.categoryRepo.Add(data)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryUsecaseImpl) UpdateCategory(id string, data *models.Category) (*models.Category, error) {
	category, err := c.categoryRepo.Update(id, data)
	if err != nil {
		return nil, err
	}
	return category, nil
}
