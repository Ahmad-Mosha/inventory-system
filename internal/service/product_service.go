package service

import (
	"fmt"

	"github.com/Ahmad-Mosha/inventory-system/internal/models"
	"github.com/Ahmad-Mosha/inventory-system/internal/repository"
)

type ProductService interface {
	RegisterProduct(p *models.Product) error
	SellProduct(id int, quantity int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) RegisterProduct(p *models.Product) error {
	if p.Price < 0 {
		return fmt.Errorf("Product price must be greater than zero")
	}

	if p.StockQuantity < 0 {
		return fmt.Errorf("business error: initial stock cannot be negative")
	}
	if p.Name == "" {
		return fmt.Errorf("business error: product name cannot be empty")
	}

	return s.repo.Create(p)
}

func (s *productService) SellProduct(id, quantity int) error {
	p, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if p.StockQuantity < quantity {
		return fmt.Errorf("Quantity is not enough")
	}
	p.StockQuantity -= quantity
	return s.repo.Update(id, p)
}
