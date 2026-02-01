package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
	"log"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll() ([]models.Category, error) {
	query := "SELECT id, name FROM categories"
	log.Printf("get query")
	rows, err := repo.db.Query(query)
	log.Printf("after query")
	if err != nil {
		log.Printf("error query: %s", err)
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var p models.Category
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, p)
	}

	return categories, nil
}

func (repo *CategoryRepository) Create(Category *models.Category) error {
	query := "INSERT INTO categories (name) VALUES ($1) RETURNING id"
	err := repo.db.QueryRow(query, Category.Name).Scan(&Category.ID)
	return err
}

// GetByID - ambil kategori by ID
func (repo *CategoryRepository) GetByID(id int) (*models.Category, error) {
	query := "SELECT id, name FROM categories WHERE id = $1"

	var p models.Category
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name)
	if err == sql.ErrNoRows {
		return nil, errors.New("kategori tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (repo *CategoryRepository) Update(Category *models.Category) error {
	query := "UPDATE categories SET name = $1 WHERE id = $2"
	result, err := repo.db.Exec(query, Category.Name, Category.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("kategori tidak ditemukan")
	}

	return nil
}

func (repo *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("kategori tidak ditemukan")
	}

	return err
}
