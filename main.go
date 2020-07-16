package main

import "fmt"

func main() {
	title := "Golang the best language"

	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("Huruf dengan index genap adalah ", string(letter))
	// 	}
	// }

	for index, letter := range title {
		if string(letter) == "a" || string(letter) == "i" || string(letter) == "u" || string(letter) == "e" || string(letter) == "o" {
			fmt.Println("Huruf dengan huruf vocal", string(letter))
		}
	}
}
