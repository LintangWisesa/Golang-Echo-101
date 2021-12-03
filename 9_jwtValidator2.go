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
type validation struct {
	Status  bool
	Message string
}

var JWT_SIGNING_METHOD = jwt.SigningMethodHS512
var JWT_SIGNATURE_KEY = []byte(cm.Config.JwtSignature)

func validate(token string) *validation {
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
			// return false
			v := validation{Status:false, Message:"invalid token: wrong signature"}
			return &v
		}
	}
	claims, ok := tkn.Claims.(*JwtClaim)
	if !ok {
		// // fmt.Println("couldn't parse claims")
		// return false
		v := validation{Status:false, Message:"invalid token: couldn't parse claims"}
		return &v
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		// fmt.Println("jwt is expired")
		// return false
		v := validation{Status:false, Message:"invalid token: token is expired"}
		return &v
	}
	// return true
	v := validation{Status:true, Message:"valid token: you're ready to go"}
	return &v
}

func main(){
	a := validate("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY5NDU2MTAsImlzcyI6Ik15c2YgUG9ja2V0IERlYWxzIiwibWRuIjoiNjI4ODg2NzE5MzI3In0.ZVwAvT3WY2qZKUEUGzbHPnYPt0xaN6JPQwtfe7eJzvZjp88v7mjKvllZYL2S-sV2SEh43WoQvI1UJLWO2A6Qfg")
	fmt.Println(a)
	fmt.Println(a.Status)
	fmt.Println(a.Message)
}