package slice

import "fmt"

func slice() {
	var friends = make([]string, 0)
	friends = append(friends, "fajar", "agus", "maulana", "zakaria", "wahyu", "zayyan", "rizki", "ramadhan", "billy", "hendar")

	for i, n := range friends {
		fmt.Println(i, n)
	}
}
