package main

import (
	"fmt"
	"strconv"
)

// package level

//  := declaration didnot work here, only full declaration will work

var x int = 5

// for declaring multiple variable we can warp it in on group of var

var (
	id   float32 = 12
	name string  = "Ramesh"
	age  int     = 23
)

var z int = 50 // shadowing

// there are three visibility in golang

var a int = 15 // at same package level only

var A int = 20 // global level

// in golang the length of variable name show its life span in program

var i int = 1  // usually used for counter purpose in loop

var EmployeeName string = "Ramesh" // can be use globally

// camelCase convention follow

var theURL string = "https://www.google.com" // for acronym keep upper case
var firstName string = "Ramesh" // usual camelcase convention

func main() {
	var a int // declaration
	a = 12    // initialization  local scope level visiblity

	var b int = 15 // short hand operator both initialization and declaration in one line, get datatype by infering 

	c := 22
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)

	d := 34.
	fmt.Printf("%T, %v", d, d) // %T for Type and %v for value

	var e int16 = 10
	fmt.Printf("\n%T, %v", e, e) // %T for Type and %v for value

	fmt.Printf("\n%T, %v", x, x) // %T for Type and %v for value

	// var y float32 = "Hello" // strictly type language => float var can store string data

	fmt.Printf("\nid = %v, name = %v, age = %v", id, name, age)

	// effect of shadowing show here
	fmt.Println("z = ", z)	// printing value of package level 
	
	var z int = 51			// initialization and declaration at local scope
	fmt.Println("z = ", z)
	// z := 52 // re-declaration not allowed
	z = 53	// assingment at local scope
	fmt.Println("z = ", z)

	fmt.Println(theURL)
	fmt.Println(firstName)

	// type conversion

	var f float32 = float32(a)
	fmt.Printf("previous type = %T type = %T, value = %v",a, f, f)

	// for string conversion we need to use one package called strconv
	var g string = strconv.Itoa(a)  // converting integer to ascii
	fmt.Printf("\nprevious type = %T type = %T, value = %v",a, g, g)

}
