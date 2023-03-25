package token

import "github.com/gin-gonic/gin"
func GetTokenHandler(c *gin.Context) {
	tokenGen, err := createToken()
	if err != nil {
		c.Data(500, "text/html", []byte(err.Error()))
		return
	}
	c.Data(200, "text/html", tokenGen)
}