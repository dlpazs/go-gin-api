package controller

import (
	"github.com/dlpazs/golang-gin-api/entity"
	"github.com/dlpazs/golang-gin-api/service"
	"github.com/gin-gonic/gin"
)

type BusinessController interface {
	FindAll() []entity.Business
	Save(ctx *gin.Context)
}

type controller struct {
	service service.BusinessService
}

func New(service service.BusinessService) BusinessController {
	return controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Business {
	return service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) {
	var business entity.Business
	ctx.BindJSON(&business)
	c.service.Save(business)
	return business
}
