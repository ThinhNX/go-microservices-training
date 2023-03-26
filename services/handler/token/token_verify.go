package token

import (
	"log"
	"os"

	"github.com/kataras/jwt"
)
func IsTokenValid(token string) bool{
	pubKey, err := os.ReadFile("public_key.pem")
	if err != nil {
		log.Default().Println("error read pubkey ", err)
		return false
	}
	pubKeyAfterParse , _ := jwt.ParsePublicKeyRSA(pubKey)
	verifiedToken , errVerify := jwt.Verify(jwt.RS256, pubKeyAfterParse, []byte(token))
	if errVerify != nil {
		log.Default().Println("error verify token ", errVerify)
		return false
	}
	var claims jwt.Claims
	errDecode := verifiedToken.Claims(&claims)
	if errDecode != nil {
		return false
	}
	if claims.ID == "admin" {
		return true
	} else {return false}
}