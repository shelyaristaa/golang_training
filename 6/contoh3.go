// package main

// import "fmt"

// type Employee struct {
// 	name string
// 	salary int
// }

// func (e *Employee) changeName(newName string) {
// 	(*e).name = newName
// }

// func main() {
// 	e := Employee{
// 		name: "Rose Geller",
// 		salary: 1200,
// 	}

// 	// e before name change
// 	fmt.Println("e before name change =", e)
// 	// create pointer to 
// }

package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width float64
	height float64
	area float64 // akan menyimpan hasil area ke dalam struct ini
  }
  
  type Circle struct {
	radius float64
  }
  
  func (r Rect) Area () float64 {
	return r.width * r.height
  }
  
  func (c Circle) Area () float64{
	return math.Pi * c.radius * c.radius
  }
  func (r *Rect) CalculateArea() {
	  r.area = r.width * r.height
  }
  func main()  {
	fmt.Println("")
	rect := Rect{
		width: 5.0,
		height: 4.0
	}
	cir := Circle {5.0}
	fmt.Printf ("Area of retangle rect = %0.2f\n", rect.Area())
	fmt.Printf ("Area of circle cir = %0.2f\n", cir.Area())

	fmt.Printf("Rectangle before call pointer method = %v\n", rect)
	rect.CalculateArea()
  }
