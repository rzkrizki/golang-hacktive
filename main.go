package main

import (
	"golang-hacktive/service"

	"github.com/gin-gonic/gin"
)

var port = ":6969"

func main() {

	router := gin.Default()
	var db []*service.User
	userSrv := service.NewUserService(db)
	router.POST("/register/:id/:name", userSrv.RegisterUser)
	router.GET("/get-user/", userSrv.GetUserRegister)
	router.GET("/getuserbyid/:id", userSrv.GetUserRegisterbyID)
	router.Run(port)
}
