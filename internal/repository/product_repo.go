package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ahmad-Mosha/inventory-system/internal/models"
)

type sqliteProductRepo struct {
	db *sql.DB
}

func NewSQLiteProductRepo(db *sql.DB) ProductRepository {
	return &sqliteProductRepo{db: db}
}

func (r *sqliteProductRepo) GetByID(id int) (*models.Product, error) {
	query := `SELECT id, name, price, stock_quantity FROM products WHERE id = ?`

	row := r.db.QueryRow(query, id)

	p := &models.Product{}

	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.StockQuantity)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *sqliteProductRepo) Create(p *models.Product) error {
	query := `INSERT INTO products (name, price, stock_quantity) VALUES (?, ?, ?)`

	_, err := r.db.Exec(query, p.Name, p.Price, p.StockQuantity)
	return err
}

func (r *sqliteProductRepo) GetAll() ([]*models.Product, error) {
	query := `SELECT id, name, price, stock_quantity FROM products`
	result, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	var products []*models.Product
	for result.Next() {
		p := &models.Product{}

		err := result.Scan(&p.ID, &p.Name, &p.Price, &p.StockQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err = result.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *sqliteProductRepo) Update(id int, p *models.Product) error {
	query := `UPDATE products SET name = ?, price = ?, stock_quantity = ? WHERE id = ?`
	res, err := r.db.Exec(query, p.Name, p.Price, p.StockQuantity, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %d not found", id)
	}

	return nil
}
func (r *sqliteProductRepo) Delete(id int) error {
	query := `DELETE FROM products WHERE id = ?`
	res, err := r.db.Exec(query, id)
	if err != nil {

		return err
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("product with ID %d not found", id)
	}
	return nil
}
