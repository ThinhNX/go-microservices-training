package handler
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func AlwaysBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "")
}