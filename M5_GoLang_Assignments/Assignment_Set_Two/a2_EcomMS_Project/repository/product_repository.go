package repository

import (
	"a2-ecomms-project/model"
	"database/sql"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) AddProduct(product model.Product) error {
	query := `INSERT INTO products (name, description, price, stock, category_id) VALUES (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, product.Name, product.Description, product.Price, product.Stock, product.CategoryID)
	return err
}

func (r *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query := `SELECT * FROM products WHERE id = ?`
	row := r.DB.QueryRow(query, id)

	var product model.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CategoryID); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) UpdateProductStock(id int, stock int) error {
	query := `UPDATE products SET stock = ? WHERE id = ?`
	_, err := r.DB.Exec(query, stock, id)
	return err
}

func (r *ProductRepository) DeleteProduct(id int) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
