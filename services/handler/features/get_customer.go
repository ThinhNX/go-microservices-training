package features

import (
	"github.com/gin-gonic/gin"
	"go-microservices-training/services/model"
)

func CustomerHandler(c *gin.Context) {
	toResponse := model.Customer{
		CustomerName: "ThinhNX",
		CustomerGender: "Male",
	}
	c.JSON(200, toResponse)
}