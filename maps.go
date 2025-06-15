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

//mutating a map
func mutating(){
var m = map[string]int{
	"one":   1,
	"two":   2,
}

fmt.Println(m)
// Insert or update an element in map m:
// m[key] = elem
 m["three"] = 3

 fmt.Println(m)
// Retrieve an element:
// elem = m[key]
	elem := m["three"]
	fmt.Println("The value of key 'three' is:", elem)

// Delete an element:
// delete(m, key)
delete(m, "three")
fmt.Println("After deletion:", m)


// Test that a key is present with a two-value assignment:
// elem, ok = m[key]
elem, ok := m["one"]
if ok {
	fmt.Println("The value of key 'one' is:", elem)
} else {
	fmt.Println("Key 'one' not found")
}

// Test that a key is present with a single-value assignment:
elem = m["two"]
fmt.Println("The value of key 'two' is:", elem)
}