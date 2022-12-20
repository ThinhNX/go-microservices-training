package router

import (
	"go-microservices-training/services/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routers struct {
	Pattern  string
	Method   string
	HandlerF func(*gin.Context)
	Name     string
}

func NewHandler() http.Handler {

	RouterList := []Routers{
		{
			Pattern:  "/customer",
			Method:   "GET",
			HandlerF: handler.CustomerHandler,
			Name:     "handlerCustomer",
		},
		{
			Pattern:  "/test",
			Method:   "GET",
			HandlerF: handler.AlwaysBadRequest,
			Name:     "handlerTest",
		},
	}
	handler := gin.New()
	gin.SetMode(gin.ReleaseMode)
	handler.SetTrustedProxies(nil)
	handlerGroup := handler.Group("/v1")
	for _, router := range RouterList {
		switch router.Method {
		case http.MethodGet:
			{
				handlerGroup.GET(router.Pattern, router.HandlerF)
			}
		case http.MethodPost:
			{
				handlerGroup.POST(router.Pattern, router.HandlerF)
			}
		case http.MethodPut:
			{
				handlerGroup.PUT(router.Pattern, router.HandlerF)
			}
		case http.MethodDelete:
			{
				handlerGroup.DELETE(router.Pattern, router.HandlerF)
			}
		}
	}
	return handler
}
