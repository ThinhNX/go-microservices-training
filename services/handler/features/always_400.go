package features

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AlwaysBadRequest(c *gin.Context) {
	Panik = !Panik
	c.JSON(http.StatusBadRequest, "")
}
