package service

import (
	"a2-ecomms-project/model"
	"a2-ecomms-project/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) AddProduct(product model.Product) error {
	return s.Repo.AddProduct(product)
}

func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.Repo.GetProductByID(id)
}

func (s *ProductService) UpdateStock(id int, stock int) error {
	return s.Repo.UpdateProductStock(id, stock)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.Repo.DeleteProduct(id)
}
