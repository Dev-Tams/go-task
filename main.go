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




    // vowels := []string{"e", "u", "a", "o", "i"}
    // sort.Strings(vowels)
    // fmt.Println(vowels)

    // for i := range 5{
    //     fmt.Println(vowels[i])
    // }

    // for index, value := range vowels{
    //     fmt.Printf("the value of %v is %v]\n", index, value)
    // }

    // greeting := greetUser("Tami")
    // fmt.Println(greeting)

    // sum := addNumbers(5, 10)
    // fmt.Printf("The sum of 5 and 10 is: %d\n", sum)

    // number := 6
    // fmt.Printf("Is %d even? %v\n", number, isEven(number))

    // number = 4
    // fmt.Printf("Is %d odd? %v\n", number, isOdd(number))



    fmt.Println("a good day for", rand.Intn(10), "hours of work")

    //working with boolean
    age := 18

    fmt.Println(age == 30)
    fmt.Println(age <= 30)
    fmt.Println(age >= 45)


    //slices
    ages := []int{19, 54, 16, 5, 56, 9, 45, 14, 89, 90}

    n := []string{"Tami", "ami", "ammy", "Tams"}


    if age <= 17 {
        fmt.Printf("age is %v \n", age)
    }
    if age <= 30 {
        fmt.Printf("age is %v \n", age)
    }


    //for loop
    for index, value := range ages{
        if value >= 18  {
            fmt.Printf("age %v is legal, continuing at position %v \n", value, index)
            continue
        }
        if value <= 18  {
            fmt.Printf("age %v is not legal, break at %v \n", value, index)
            break
        }
    }

    //sort ages
     sort.Ints(ages)
     fmt.Println("inputed ages are", ages)

     chatToUsers(n, greetUser)

     //loop over food map
     menu()
     
       
}
// Function to greet a user



func chatToUsers(n []string, f func(string) ){
    for _, value := range n{
        f(value)
    }
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




