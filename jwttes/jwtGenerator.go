package main

import (
	"fmt"
	"time"
	"strconv"
	jwt "github.com/golang-jwt/jwt"
	cm "GoLang-echo-101/common"
)

type JwtClaim struct {
    jwt.StandardClaims
    Mdn string `json:"mdn"`
}

var JWT_SIGNING_METHOD = jwt.SigningMethodHS512
var JWT_SIGNATURE_KEY = []byte(cm.Config.JwtSignature)

func CanonicalizeMdn(mdn string) string {
	if mdn[0:1] == "+" {
		mdn = mdn[1:]
	} else if mdn[0:1] == "0" {
		mdn = "62" + mdn[1:]
	} else if !(mdn[0:2] == "62") {
		mdn = "62" + mdn
	}
	return mdn
}

func generate(mdn string) string {
	
	// userInfo := map[string]string{"mdn": "628886719327"}
	// fmt.Println(userInfo["mdn"])
	
	minutes, err := strconv.Atoi(cm.Config.JwtMinutesExpired)
	claim := JwtClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer:    cm.Config.JwtIssuer,
			ExpiresAt: time.Now().Add(time.Duration(minutes) * time.Minute).Unix(),
		},
		Mdn: CanonicalizeMdn(mdn),
	}
	
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claim,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "error"
	}
	return signedToken
}

func main(){
	fmt.Println(generate("628886719327"))
}