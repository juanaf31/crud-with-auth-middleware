package productrepository

import (
	"database/sql"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/utils"
)

type ProductRepoImpl struct {
	db *sql.DB
}

func (s *ProductRepoImpl) GetAll() ([]*models.Product, error) {
	rows, err := s.db.Query(utils.GET_ALL_PRODUCTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listProducts []*models.Product
	for rows.Next() {
		products := models.Product{}
		err := rows.Scan(&products.ID, &products.ProductCode, &products.ProductName, &products.ProductCategory.CategoryId, &products.ProductCategory.CategoryName)
		if err != nil {
			return nil, err
		}
		listProducts = append(listProducts, &products)
	}
	return listProducts, nil

}

func (s *ProductRepoImpl) GetByID(id string) (*models.Product, error) {

	row := s.db.QueryRow(utils.GET_PRODUCT_BY_ID, id)
	var product = models.Product{}
	err := row.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.ProductCategory.CategoryId, &product.ProductCategory.CategoryName)
	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (s *ProductRepoImpl) Delete(id string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(utils.DELETE_PRODUCT, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}

func (s *ProductRepoImpl) Add(products []*models.Product) ([]*models.Product, error) {
	for _, product := range products {

		tx, err := s.db.Begin()
		if err != nil {
			return nil, err
		}
		res, err := tx.Exec(utils.ADD_CATEGORY, product.ProductCategory.CategoryId, product.ProductCategory.CategoryName, utils.CurDate, utils.CurDate)

		if err != nil {
			tx.Rollback()
			return nil, err
		}
		id, err := res.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		_, err = tx.Exec(utils.ADD_PRODUCT, product.ID, product.ProductCode, product.ProductName, id, utils.CurDate, utils.CurDate)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		return nil, tx.Commit()
	}

	return products, nil
}

func (s *ProductRepoImpl) Update(id string, product *models.Product) (*models.Product, error) {
	tx, err := s.db.Begin()
	_, err = tx.Exec(utils.UPDATE_PRODUCT, product.ID, product.ProductCode, product.ProductName, product.ProductCategory.CategoryId, utils.CurDate, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return product, tx.Commit()
}

func InitProductRepoImpl(db *sql.DB) ProductRepository {
	return &ProductRepoImpl{db: db}
}
