package controller

import (
	"net/http"
	"product-manager/model"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	// Usecase
}

func NewProductController() ProductController {
	return ProductController{}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "batatinha",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
