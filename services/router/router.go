package router

import (
	"go-microservices-training/services/handler/features"
	"go-microservices-training/services/handler/login"

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
		HandlerF: features.CustomerHandler,
		Name:     "featuresCustomer",
	},
	{
		Pattern:  "/test",
		Method:   "GET",
		HandlerF: features.AlwaysBadRequest,
		Name:     "featuresTest",
	},
	{
		Pattern:  "shoplist",
		Method:   "GET",
		HandlerF: features.GetShopHandler,
		Name:     "handleShopListGet",
	},
	{
		Pattern:  "login",
		Method:   "POST",
		HandlerF: login.LoginHandler,
		Name:     "handleShopListGet",
	},
}

func NewHandler() http.Handler {
	features := gin.New()
	gin.SetMode(gin.DebugMode)
	features.SetTrustedProxies(nil)
	featuresGroup := features.Group("/v1")
	for _, router := range routerList {
		switch router.Method {
		case http.MethodGet:
			{
				featuresGroup.GET(router.Pattern, router.HandlerF)
			}
		case http.MethodPost:
			{
				featuresGroup.POST(router.Pattern, router.HandlerF)
			}
		case http.MethodPut:
			{
				featuresGroup.PUT(router.Pattern, router.HandlerF)
			}
		case http.MethodDelete:
			{
				featuresGroup.DELETE(router.Pattern, router.HandlerF)
			}
		}
	}
	return features
}
