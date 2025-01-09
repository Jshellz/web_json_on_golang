package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/db/connection"
	"web/db/models"
)

type CreateProductInput struct {
	Title string  `json:"title" binding:"required"`
	Price float32 `json:"price" binding:"required"`
	Photo string  `json:"photo" binding:"required"`
}

func FindAllProducts(ctx *gin.Context) {
	var product []models.Product
	connection.DB.Find(&product)

	ctx.JSON(http.StatusOK, gin.H{
		"products": product,
	})
}

func CreateProduct(ctx *gin.Context) {
	var createProduct CreateProductInput
	if err := ctx.ShouldBindJSON(&createProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	product := models.Product{Title: createProduct.Title, Price: createProduct.Price, Photo: createProduct.Photo}
	connection.DB.Create(&product)

	ctx.JSON(http.StatusCreated, gin.H{
		"product": product,
	})
}

func FindProductByID(ctx *gin.Context) {
	var product models.Product

	if err := connection.DB.Where("id = ?", ctx.Param("id")).First(&product).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "Record not found!",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": &product,
	})
}
