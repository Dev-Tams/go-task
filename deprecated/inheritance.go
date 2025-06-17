package deprecated

import (
	// "io"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

//reader interface
//writer interface
//closer interface


// func doSomething(r io.Writer )error {
// 	// do something with the reader
// 	// do something with the writer
// 	// do something with the closer
// 	return nil	
// }
type rect struct {
	width, height float64
}


type circle struct {
	radius float64
}


//rect
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}


//circle

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}


var r = rect{
	width: 5,
	height: 10,
  }
 
 
var c = circle{
	radius: 30,
}



