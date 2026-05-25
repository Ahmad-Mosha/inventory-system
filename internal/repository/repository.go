package repository

import "github.com/Ahmad-Mosha/inventory-system/internal/models"


type ProductRepository interface {
    Create(p *models.Product) error
    GetByID(id int) (*models.Product, error)
    GetAll() ([]*models.Product, error) 
    Update(id int, p *models.Product) error
    Delete(id int) error
}

type CustomerRepository interface {
    Create(c *models.Customer) error
    GetByID(id int) (*models.Customer, error)
    GetAll() ([]*models.Product, error) 
    Update(id int, c *models.Customer) error
    Delete(id int) error
}

type OrderRepository interface {
    Create(o *models.Order) error
    GetByID(id int) (*models.Order, error)
    GetByCustomerID(customerID int) ([]*models.Order, error)
}