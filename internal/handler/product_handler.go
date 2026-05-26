package handler

import (
	"github.com/Ahmad-Mosha/inventory-system/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	srv service.ProductService
}

func NewProductHandler(srv service.ProductService) *ProductHandler {
	return &ProductHandler{srv: srv}
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	
}
