package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friends struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func assignemt1() {

	args := os.Args[1]
	i, _ := strconv.Atoi(args)

	getFriend := getFriend(i)
	fmt.Println(getFriend)
}

func getFriend(number int) Friends {
	lists := []Friends{
		{
			nama:      "fajar",
			alamat:    "bonsir",
			pekerjaan: "developer",
			alasan:    "belajar",
		},
		{
			nama:      "rizki",
			alamat:    "bonsir",
			pekerjaan: "developer",
			alasan:    "belajar",
		},
		{
			nama:      "zaka",
			alamat:    "bonsir",
			pekerjaan: "developer",
			alasan:    "belajar",
		},
	}

	return lists[number]
}
