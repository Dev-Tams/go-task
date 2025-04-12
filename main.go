package main

import "fmt"

func main() {

    task := "learning go!"
    name := "Tami"
    fmt.Println("Hello, World!")
    
    fmt.Println("My name is ", name, "and i took on a new task;", task)
    fmt.Printf("My name is %v and i took on a new task; %v \n", name, task)

    var str = fmt.Sprintf("My name is %v and i took on a new task; %v", name, task)

    fmt.Print(str)


}