package token

import (
	"go-microservices-training/services/model"
	"io/ioutil"
	"log"
	"time"

	"github.com/kataras/jwt"
)

func createToken() ([]byte, error) {
	// pubKey, _ := os.ReadFile("public_key.pem")
	prvKey, errRead := ioutil.ReadFile("private_key.pem")
	if errRead != nil {
		log.Default().Println("read key err: ", errRead)
		return nil, errRead
	}
	prvKeyAfterParse, _ := jwt.ParsePrivateKeyRSA(prvKey)

	payload := model.UserClaims{
		UserID:     "id thu nghie",
		UserWaller: "1 ty dong",
	}
	token, err := jwt.Sign(jwt.RS256, prvKeyAfterParse, payload, jwt.MaxAge(100*time.Hour))
	if err != nil {
		log.Default().Println("Sign token err: ", err)

		return nil, err
	}
	return token, nil
}
