package main

import (
	"fmt"
	"golang-hacktive/service"
)

func main() {
	userSrv := service.NewUserService()
	res := userSrv.Register(&service.User{Nama: "budi"})
	fmt.Println(res)
}
