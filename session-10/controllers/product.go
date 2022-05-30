package controllers

import (
	"fmt"
	"sesi10/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		db: db,
	}
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	var product = models.Product{}

	var err = c.ShouldBind(&product)
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}

	id := c.MustGet("id").(float64)

	product.UserID = uint(id)

	err = p.db.Create(&product).Error
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"payload": product,
	})
}

func (p *ProductController) GetProducts(c *gin.Context) {
	var products []models.Product

	err := p.db.Model(models.Product{}).Joins("JOIN users ON users.id = products.user_id").Scan(&products).Error
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"payload": products,
	})
}
