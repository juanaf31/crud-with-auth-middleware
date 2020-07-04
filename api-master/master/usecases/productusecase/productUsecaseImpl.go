package productusecase

import (
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/api-master/master/repositories/productrepository"
)

type ProductUsecaseImpl struct {
	productRepo productrepository.ProductRepository
}

func InitProductUsecase(productRepo productrepository.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{productRepo: productRepo}
}

func (p *ProductUsecaseImpl) GetProducts() ([]*models.Product, error) {
	products, err := p.productRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductUsecaseImpl) GetProductByID(id string) (*models.Product, error) {
	product, err := p.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductUsecaseImpl) DeleteProduct(id string) error {
	err := s.productRepo.Delete(id)
	return err
}
func (s *ProductUsecaseImpl) AddProduct(data []*models.Product) ([]*models.Product, error) {
	product, err := s.productRepo.Add(data)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductUsecaseImpl) UpdateProduct(id string, data *models.Product) (*models.Product, error) {
	teacher, err := s.productRepo.Update(id, data)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}
