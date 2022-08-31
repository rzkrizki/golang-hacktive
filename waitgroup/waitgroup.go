package main

import (
	"fmt"
	"golang-hacktive/service"
	"sync"
)

func WaitGroup() {
	var db []*service.User
	userSrv := service.NewUserService(db)
	user1 := userSrv.Register(&service.User{ID: 1, Nama: "rizki"})
	fmt.Println(user1)
	user2 := userSrv.Register(&service.User{ID: 2, Nama: "zaka"})
	fmt.Println(user2)
	user3 := userSrv.Register(&service.User{ID: 3, Nama: "fajar"})
	fmt.Println(user3)
	fmt.Println("-----------Hasil get user-------------")
	resGet := userSrv.GetUser()

	var wg sync.WaitGroup
	wg.Add(len(resGet))

	for _, v := range resGet {
		go cetakNama(&wg, v.Nama)
	}

	wg.Wait()
	// fmt.Println("-----------Hasil get user by id-------------")
	// getusr1 := userSrv.GetUserbyID(1)
	// fmt.Println(getusr1)
}

func cetakNama(wg *sync.WaitGroup, nama string) {
	fmt.Println(nama)
	wg.Done()
}
