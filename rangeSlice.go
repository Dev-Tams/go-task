package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

var opp = []int{}

// rangeSlice demonstrates the use of range to iterate over slices.
// It prints the powers of 2 and multiples of 3.
func rangeSlice() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	
	
	fmt.Println(opp)
	
	opp = append(opp, 0, 3, 6, 9, 12, 15, 18, 21)

	for i, v := range opp {
		fmt.Printf("3*%d = %d\n", i, v)
	}
}
