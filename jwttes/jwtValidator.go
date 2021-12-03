package main

import (
	"fmt"
	"time"
	jwt "github.com/golang-jwt/jwt"
	cm "GoLang-echo-101/common"
)

type JwtClaim struct {
    jwt.StandardClaims
    Mdn string `json:"mdn"`
}
var JWT_SIGNING_METHOD = jwt.SigningMethodHS512
var JWT_SIGNATURE_KEY = []byte(cm.Config.JwtSignature)

func validate(token string) bool {
	
	// token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJtZG4iOiI2Mjg4ODY3MTkzMjcifQ.Z6z7_APGYmNjyP_UZJkkKAmd0IpCgPFu0aD4O2RzPNUV8wizJoyGg_7vcVy6VDKx2eIIo7VRFdoL0RZ-zUT-5w"
	
	claims := &JwtClaim{}

	tkn, err := jwt.ParseWithClaims(
		token,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return JWT_SIGNATURE_KEY, nil 
		},
	)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// fmt.Println("Invalid signature")
			return false
		}
	}

	claims, ok := tkn.Claims.(*JwtClaim)
	if !ok {
		// fmt.Println("couldn't parse claims")
		return false
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		// fmt.Println("jwt is expired")
		return false
	}

	// fmt.Println(claims)
	// fmt.Println(claims.Mdn)
	// fmt.Println(time.Now().UTC().Unix())
	// fmt.Println(claims.ExpiresAt)
	// fmt.Println(claims.Issuer)
	// return claims.Mdn
	return true
}

func main(){
	fmt.Println(validate("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY1MzczMDcsImlzcyI6Ik15U21hcnRmcmVuIiwibWRuIjoiNjI4ODg2NzE5MzI3In0.yHYzH1dydxjSj0XGFMkZ27BWTiNaSZIX84u6kYe92wdYDF2LWHMM7GzBFFnbxcmMP9nW3hfOXtevTt_VGqHPCQ"))
}