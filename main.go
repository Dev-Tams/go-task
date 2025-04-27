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
	// final := concat(1, 2, 4)
	// fmt.Println(final)
// 
	//  ar := checkMail(9)

	// ev := &EmailValidator{Email: "tammyalemu@gmail.com"}
	// countLength(ev)
	// ev.checkMail()

	carModel :=  model{"red", "sport", 224, brand{"Toyota", "suv", 2009}}

	fmt.Println(carModel.color, carModel.engineNo, carModel.brand)



	
    // names := []string{"bob", "sue", "alice"}
    // printStrings(names...)

    // harryPotterAggregator := concatter()
	// harryPotterAggregator("Mr.")
	// harryPotterAggregator("and")
	// harryPotterAggregator("Mrs.")
	// harryPotterAggregator("Dursley")
	// harryPotterAggregator("of")
	// harryPotterAggregator("number")
	// harryPotterAggregator("four,")
	// harryPotterAggregator("Privet")

	// fmt.Println(harryPotterAggregator("Drive"))

    // r := doSomething(nil)
    // println(r)
    // r is nil, so it will not print anything
}


func printStrings(strings ...string) {
    for i := 0; i < len(strings); i++ {
        fmt.Println(strings[i])
    }

}

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}


