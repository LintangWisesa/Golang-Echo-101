package main

import (
	"fmt"
	cm "GoLang-echo-101/common"
)

func main(){
	fmt.Println(cm.Config.JwtSignature)
	fmt.Sprint(cm.Config.JwtMinutesExpired)
	fmt.Println(CanonicalizeMdn("08886719327"))
}