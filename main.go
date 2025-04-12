package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {


    // working with srings 

    // task := "learning go!"
    // name := "Tami"
    // fmt.Println("Hello, World!")
    
    // fmt.Println("My name is ", name, "and i took on a new task;", task)
    // fmt.Printf("My name is %v and i took on a new task; %v \n", name, task)

    // var str = fmt.Sprintf("My name is %v and i took on a new task; %v", name, task)

    // fmt.Print(str)

    
    // // arrrays

    // vowels := [5] string{"a", "e", "i", "o", "u"}

    // fmt.Println(vowels, len(vowels))


    // //slices 
    // consonant := [] int {21}

    // fmt.Println(consonant)

    // //ranges

    // rangeOne := vowels[2:5]
    // rangeTwo := vowels[:4]
    // rangeThree := vowels[:5]

    // rangeTwo = append(rangeTwo, "vowels")
    // fmt.Println(rangeOne, rangeTwo, rangeThree)


    //slices
    ages := []int{12, 54, 24, 5, 56, 23, 45, 67, 89, 90}
    sort.Ints(ages)
    fmt.Println(ages)


    vowels := []string{"e", "u", "a", "o", "i"}
    sort.Strings(vowels)
    fmt.Println(vowels)

 

    greeting := greetUser("Tami")
    fmt.Println(greeting)

    sum := addNumbers(5, 10)
    fmt.Printf("The sum of 5 and 10 is: %d\n", sum)

    number := 6
    fmt.Printf("Is %d even? %v\n", number, isEven(number))

    number = 4
    fmt.Printf("Is %d odd? %v\n", number, isOdd(number))

    fmt.Println("a good day for", rand.Intn(10), "hours of work")


}
// Function to greet a user
func greetUser(name string) string {
    return fmt.Sprintf("Welcome, %s!", name)
}

// Function to calculate the sum of two numbers
func addNumbers(a int, b int) int {
    return a + b
}

// Function to check if a number is even
func isEven(number int) bool {
    return number%2 == 0
}

func isOdd(number int) bool {
    return number%2 != 0
}



