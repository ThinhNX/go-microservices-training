package login

import (
	"go-microservices-training/services/model"
	"log"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	// jsd := json.NewDecoder(c.Request.Body)
	// var loginBody model.User
	// err := jsd.Decode(&loginBody)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	c.Data(http.StatusUnauthorized, "", []byte(err.Error()))
	// }
	queryP := c.Request.URL.Query()
	userName := queryP.Get("user")
	pass := queryP.Get("pass")
	log.Println("Login User: ", userName)
	for _, v := range model.UserList {
		if v.Name == userName && v.Password == pass {
			c.Data(200, "text/html", nil)
		} else {
			c.Data(401, "text/html", nil)
		}
	}
}
