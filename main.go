package main

import (
	_ "baiduOcr4go/routers"

	"github.com/astaxie/beego"
)

func main() {
	// var Token = actions.GetToken()
	// if Token.Error == "" {
	// 	fmt.Println("Has Token")
	// 	fmt.Println(Token)
	// }

	beego.Run()
}
