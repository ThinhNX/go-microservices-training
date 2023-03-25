package login

import (
	"encoding/json"
	"go-microservices-training/services/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	jsd := json.NewDecoder(c.Request.Body)
	var loginBody model.User
	err := jsd.Decode(&loginBody)
	if err != nil {
		c.Data(http.StatusUnauthorized, "", []byte(err.Error()))
	}
	for _, v := range model.UserList {
		if v.Name == loginBody.Name && v.Password == loginBody.Password {
			c.Data(200, "text/html", nil)
		} else {
			c.Data(401, "text/html", nil)
		}
	}
}
