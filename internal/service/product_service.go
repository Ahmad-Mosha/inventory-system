package service

import (
	"fmt"

	"github.com/Ahmad-Mosha/inventory-system/internal/models"
	"github.com/Ahmad-Mosha/inventory-system/internal/repository"
)

type ProductService interface {
	GetAll() ([]*models.Product, error)
	GetByID(id int) (*models.Product, error)
	Update(id int, p *models.Product) error
	Delete(id int) error
	SellProduct(id int, quantity int) error
	RegisterProduct(p *models.Product) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll() ([]*models.Product, error) {
	result, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *productService) GetByID(id int) (*models.Product, error) {
	result, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *productService) Update(id int, p *models.Product) error {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if p.Name != "" {
		existing.Name = p.Name
	}

	if p.Price > 0 {
		existing.Price = p.Price
	}

	if p.StockQuantity > 0 {
		existing.StockQuantity = p.StockQuantity
	}

	return s.repo.Update(id, existing)
}
func (s *productService) Delete(id int) error {
	return s.repo.Delete(id)
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
