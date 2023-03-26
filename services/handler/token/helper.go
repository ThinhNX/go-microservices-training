package token

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
)

func createToken(c *gin.Context) ([]byte, error) {
	// pubKey, _ := os.ReadFile("public_key.pem")
	prvKey, errRead := ioutil.ReadFile("private_key.pem")
	if errRead != nil {
		log.Default().Println("read key err: ", errRead)
		return nil, errRead
	}
	prvKeyAfterParse, _ := jwt.ParsePrivateKeyRSA(prvKey)
	errParseForm := c.Request.ParseForm()
	if errParseForm != nil {
		return nil, errParseForm
	}
	idRequest := c.PostForm("ID")
	tokenClaims := jwt.Claims{
		ID: idRequest,
	}
	token, err := jwt.Sign(jwt.RS256, prvKeyAfterParse, tokenClaims, jwt.MaxAge(100*time.Hour))
	if err != nil {
		log.Default().Println("Sign token err: ", err)

		return nil, err
	}
	return token, nil
}
