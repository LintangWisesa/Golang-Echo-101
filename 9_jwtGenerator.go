package main

import (
	"fmt"
	"time"
	jwt "github.com/golang-jwt/jwt"
)

type JwtClaim struct {
    jwt.StandardClaims
    Mdn string `json:"mdn"`
}
var JWT_SIGNING_METHOD = jwt.SigningMethodHS512
var JWT_SIGNATURE_KEY = []byte("321@nerftramSyM")

func main(){
	userInfo := map[string]string{"mdn": "628886719327"}
	fmt.Println(userInfo["mdn"])
	
	claim := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "MySmartfren",
			ExpiresAt: time.Now().Add(time.Duration(60) * time.Minute).Unix(),
		},
		Mdn: userInfo["mdn"],
	}
	
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claim,
	)
	
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	fmt.Println(signedToken)
	if err != nil {
		fmt.Println("Error")
	}
}