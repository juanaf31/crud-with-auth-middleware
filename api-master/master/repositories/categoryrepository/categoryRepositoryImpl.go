package categoryrepository

import (
	"database/sql"
	"liveCodeAPI/api-master/master/models"
	"liveCodeAPI/utils"
)

type CategoryRepoImpl struct {
	db *sql.DB
}

func (s *CategoryRepoImpl) GetAll() ([]*models.Category, error) {
	query := `select id, category_name from m_category`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listCategory []*models.Category
	for rows.Next() {
		category := models.Category{}
		err := rows.Scan(&category.CategoryId, &category.CategoryName)
		if err != nil {
			return nil, err
		}
		listCategory = append(listCategory, &category)
	}
	return listCategory, nil

}

func (s *CategoryRepoImpl) GetByID(id string) (*models.Category, error) {

	query := `select id, category_name from m_category where id=?`
	row := s.db.QueryRow(query, id)
	var category = models.Category{}
	err := row.Scan(&category.CategoryId, &category.CategoryName)
	if err != nil {
		return nil, err
	}
	return &category, nil

}

func (s *CategoryRepoImpl) Delete(id string) error {
	query := "delete from m_category where id=?"
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}

func (s *CategoryRepoImpl) Add(categories []*models.Category) ([]*models.Category, error) {

	for _, category := range categories {

		tx, err := s.db.Begin()
		if err != nil {
			return nil, err
		}
		_, err = tx.Exec(`insert into m_category(id,category_name,created_at,updated_at) values(?,?,?,?)`, &category.CategoryId, &category.CategoryName, utils.CurDate, utils.CurDate)

		if err != nil {
			tx.Rollback()
			return nil, err
		}

		return nil, tx.Commit()
	}

	return categories, nil
}

func (s *CategoryRepoImpl) Update(id string, category *models.Category) (*models.Category, error) {

	tx, err := s.db.Begin()
	_, err = tx.Exec("update m_category set id=?, category_name=?, updated_at=? where id=?", category.CategoryId, category.CategoryName, utils.CurDate, utils.CurDate)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return category, tx.Commit()
}

func InitCategoryRepoImpl(db *sql.DB) CategoryRepository {
	return &CategoryRepoImpl{db: db}
}
