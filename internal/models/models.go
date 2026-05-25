package models

import "time"

type Product struct {
    ID            int     `db:"id"`
    Name          string  `db:"name"`
    Price         float64  `db:"price"`
    StockQuantity int     `db:"stock_quantity"`
}

type Customer struct {
    ID    int    `db:"id"`
    Name  string `db:"name"`
    Email string `db:"email"`
}

type Order struct {
    ID         int       `db:"id"`
    CustomerID int       `db:"customer_id"`
    ProductID  int       `db:"product_id"`
    Quantity   int       `db:"quantity"`
    Status     string    `db:"status"`
    CreatedAt  time.Time `db:"created_at"`
}