package main

import "fmt"

func concat(strs ...int) int {
	final := 3
	// strs is just a slice of strings
	for str := range strs {
		final += str
	}
	return final
}

func main() {
	final := concat(1, 2, 4)
	fmt.Println(final)


    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}
func printStrings(strings ...string) {
    for i := 0; i < len(strings); i++ {
        fmt.Println(strings[i])
    }
}
