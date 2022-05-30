package controllers

import (
	"fmt"
	"net/http"
	"sesi10/helper"
	"sesi10/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (u *UserController) UserRegister(c *gin.Context) {

	user := models.User{}

	err := c.ShouldBind(&user)

	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	err = u.db.Create(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(400, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"payload": user,
	})
}

func (u *UserController) UserLogin(c *gin.Context) {
	var user models.User

	err := c.ShouldBind(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	var userDB models.User

	err = u.db.Where("email=?", user.Email).Take(&userDB).Error
	if err != nil {
		c.AbortWithStatusJSON(400, err)
		return
	}

	isOk := helper.ComparePass(userDB.Password, user.Password)
	if !isOk {
		c.AbortWithStatusJSON(401, fmt.Errorf("invalid email / password"))
		return
	}

	token := helper.GenerateToken(userDB.ID, userDB.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
