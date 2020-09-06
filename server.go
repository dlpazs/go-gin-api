package main

import (
	"github.com/dlpazs/go-gin-api/controller"
	"github.com/dlpazs/go-gin-api/service"
	"github.com/gin-gonic/gin"
)

var (
	businessService    service.BusinessService       = service.New()
	businessController controller.BusinessController = controller.New(businessService)
)

func main() {
	server := gin.Default()

	server.GET("/busingo run maess", func(ctx *gin.Context) {
		ctx.JSON(200, businessController.FindAll())
	})

	server.POST("/business", func(ctx *gin.Context) {
		ctx.JSON(200, businessController.Save(ctx))
	})

	server.Run(":8080")

}
