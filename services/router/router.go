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

var routerList = []Routers{
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
	{
		Pattern:  "shoplist",
		Method:   "GET",
		HandlerF: handler.GetShopHandler,
		Name:     "handleShopListGet",
	},
}

func NewHandler() http.Handler {
	handler := gin.New()
	gin.SetMode(gin.DebugMode)
	handler.SetTrustedProxies(nil)
	handlerGroup := handler.Group("/v1")
	for _, router := range routerList {
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
