package main

import "fmt"

func menu() {
	food := map[int]string{
		1: "rice",
		2: "bread",
		3: "beans",
		4: "salad",
	}

	food[3] = "Yam"
	for k, v := range food {
		fmt.Println(k, "-", v)
	}

	email := 1

	if length := getLength(email); length < 3 {
		fmt.Println("Email is invalid")
	}

}

func getLength(e int) int{
	return e

}



