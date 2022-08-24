package main

import "fmt"

type Friends struct {
	nama string
}

func main() {

	lists := []*Friends{
		{nama: "fajar"},
		{nama: "agus"},
		{nama: "maulana"},
		{nama: "zakaria"},
		{nama: "wahyu"},
		{nama: "zayyan"},
		{nama: "rizki"},
		{nama: "ramadhan"},
		{nama: "billy"},
		{nama: "hendar"},
	}

	var cetak = func(lists []*Friends) []string {

		var results []string

		for _, n := range lists {
			results = append(results, n.nama)
		}

		return results
	}

	fmt.Println(cetak(lists))
}
