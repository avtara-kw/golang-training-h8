package controllers

import (
	"net/http"
	"project-sesi8/params"
	"project-sesi8/services"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	personService services.PersonService
}

func NewPersonController(service *services.PersonService) *PersonController {
	return &PersonController{
		personService: *service,
	}
}

func (p *PersonController) CreateNewPerson(c *gin.Context) {
	var req params.CreatePerson

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}

	response := p.personService.CreatePerson(req)
	c.JSON(response.Status, response)
}

func (p *PersonController) GetAllPerson(c *gin.Context) {
	response := p.personService.GetAllPerson()
	c.JSON(response.Status, response)
}
