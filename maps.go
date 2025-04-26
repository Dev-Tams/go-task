package main

import (
	"fmt"
)

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

}
