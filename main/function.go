package main

import "fmt"

func main() {

	var friends = join()

	for i, n := range friends {
		fmt.Println(i, n)
	}
}

func join() []string {
	var friends = make([]string, 0)
	friends = append(friends, "fajar", "agus", "maulana", "zakaria", "wahyu", "zayyan", "rizki", "ramadhan", "billy", "hendar")

	return friends
}
