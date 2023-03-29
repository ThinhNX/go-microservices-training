package features

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Panik bool

func PanicSv(c *gin.Context) {
	defer PanicRecover(c)
	panic("panic already")
}
func makeFailOverNoti(c *gin.Context) {
	requestUrl := "http://localhost:8080/v1/panicHandling"
	res, err := http.Get(requestUrl)
	if err != nil {
		log.Default().Panic(err)
	}
	log.Default().Println("response: ", res)
	c.Status(500)
}
func PanicRecover(c *gin.Context) {
	if r := recover(); r != nil {
		log.Default().Println("panic recover from: ", r)
		makeFailOverNoti(c)
	}
}
func GonnaPanicOrNot(c *gin.Context) {
	c.JSON(200, "panic test passed")
}
func PanikCheckMdw() gin.HandlerFunc {
	return func(c *gin.Context) {
		Panik = !Panik
		log.Default().Printf("Request from %s: %s", c.ClientIP(), c.Request.URL.Path)
		log.Default().Println(" panic: ", Panik)
		if Panik {
			c.AbortWithStatusJSON(500, "panic already in panic check")
		}
		c.Next()
	}
}
