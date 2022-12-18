package router
import (
	"go-microservices-training/services/handler"
	"github.com/gin-gonic/gin"
	"log"
)
func StartSv(port string) {
	log.Println("starting sv at localhost, port: ", port)
	handlerGroup(port)
}

func handlerGroup(port string) {
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost"})
	v1Api := router.Group("/v1")
	v1Api.GET("/customer", handler.CustomerHandler)
	v1Api.GET("/test", handler.AlwaysBadRequest)
	addr := "localhost:" + port
	router.Run(addr)

}