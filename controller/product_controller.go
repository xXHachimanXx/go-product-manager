package controller

import (
	"net/http"
	"product-manager/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(productUsecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: productUsecase,
	}
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProductUseCase()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
