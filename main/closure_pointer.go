package main

import (
	"fmt"
)

func closure_pointer() {
	fajar := "fajar"
	agus := "agus"
	maulana := "maulana"
	zakaria := "zakaria"
	wahyu := "wahyu"
	zayyan := "zayyan"
	rizki := "rizki"
	ramadhan := "ramadhan"
	billy := "billy"
	hendar := "hendar"

	lists := []*string{&fajar, &agus, &maulana, &zakaria, &wahyu, &zayyan, &rizki, &ramadhan, &billy, &hendar}

	var cetak = func(lists []*string) []string {

		var results []string

		for _, n := range lists {
			results = append(results, *n)
		}

		return results
	}

	fmt.Println(cetak(lists))
}
