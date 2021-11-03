package main

import (
	"fmt"
	"math"
)

const a int = 12

const (
	e = iota
	f = iota
	g = iota
)

func main() {

	// const FirstConstant int = 12 // constant for global export
	const firstConstant int = 12 // constant for local scope, once assign can't be change, computed at compile time
	fmt.Printf("%v, %T\n", firstConstant, firstConstant)

	// firstConstant = 13  // once assign can't be change
	// fmt.Printf("%v, %T\n", firstConstant, firstConstant)

	// const secondConstant float64 = math.Tan(3.2)  // here math.Tan can't be calculated at compile time
	// fmt.Printf("%v, %T\n", secondConstant, secondConstant)

	var secondConstant float64 = math.Tan(3.2) // here math.Tan can be calculated at run time
	fmt.Printf("%v, %T\n", secondConstant, secondConstant)

	// const a int = 12
	const b float64 = 3.14
	const c string = "Ram"
	const d bool = true

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

	// arrays can't be used as const

	fmt.Printf("value of at package level %v\n", a)
	const a int16 = 16
	fmt.Printf("value of at local level %v\n", a)

	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", g)
}
