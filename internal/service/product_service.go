package service

import (
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

func (s *productService) RegisterProduct(p *models.Product) error
func (s *productService) SellProduct(id, quantity int) error
