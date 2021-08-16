package main

import (
	//"GoCode/Controller"
	"GoCode/Controller"
	"GoCode/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	db := common.InitDB()
	defer db.Close()

	REngin := gin.Default()
	REngin.Static("/assets", "./assets")
	REngin.LoadHTMLGlob("views/*")
	registRouter(REngin)

	REngin.Run(":9999")

}
func registRouter(r *gin.Engine) {
	// new(Controller.TestController).Router(r)
	// new(Controller.UserController).Router(r)
	new(Controller.PubController).Router(r)

}
