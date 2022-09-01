package main

import (
	"net/http"
)

var port = ":6969"

func WebService() {
	// var db []*service.User
	// userSrv := service.NewUserService(db)
	// http.HandleFunc("/register", userSrv.RegisterUser)
	// http.HandleFunc("/get-user/", userSrv.GetUserRegister)
	// http.HandleFunc("/getuserbyid", userSrv.GetUserRegisterbyID)
	http.ListenAndServe(port, nil)
}
