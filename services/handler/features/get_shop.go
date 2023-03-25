package features

import (
	"go-microservices-training/services/model"

	"github.com/gin-gonic/gin"
)

func GetShopHandler(c *gin.Context) {
	rspShop := model.ShopList{
		Address:  "12 hang bai",
		ShopName: "aha",
	}
	c.JSON(200, rspShop)
}
