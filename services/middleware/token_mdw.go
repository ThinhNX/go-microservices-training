package middleware

import (
	"go-microservices-training/services/handler/token"
	"log"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Default().Printf("Request from %s: %s", c.ClientIP(), c.Request.URL.Path)
		tokenFHeader := c.Request.Header.Get("Authorization")
		log.Default().Println("Token: ", tokenFHeader)
		if !token.IsTokenValid(tokenFHeader) {
			log.Default().Println("abort")
			c.AbortWithStatusJSON(401, "U have no access right")
		}
		c.Next()
	}
}
