package main

import (
	"fmt"
	"golang-hacktive/service"
)

func main() {
	var db []*service.User
	userSrv := service.NewUserService(db)
	user1 := userSrv.Register(&service.User{Nama: "rizki"})
	fmt.Println(user1)
	user2 := userSrv.Register(&service.User{Nama: "zaka"})
	fmt.Println(user2)
	user3 := userSrv.Register(&service.User{Nama: "fajar"})
	fmt.Println(user3)
	fmt.Println("-----------Hasil get user-------------")
	resGet := userSrv.GetUser()
	for _, v := range resGet {
		fmt.Println(v.Nama)
	}
}
