package main

import (
	"fmt"
)

func main() {

	var n bool = false

	fmt.Printf("%v, %T\n", n, n)

	var m int8 = 125
	fmt.Printf("%v, %T\n", m, m)

	a := 10
	b := 3
	fmt.Printf("%v\n", a+b)
	fmt.Printf("%v\n", a-b)
	fmt.Printf("%v\n", a*b)
	fmt.Printf("%v\n", a/b)
	fmt.Printf("%v\n", a%b)

	// for performing arithmetic operations we need to use same type of data
	// var m int8 = 125
	// var n int = 15
	// fmt.Printf("%v\n", a+b)  // invalid operation throw compile time error

	// bitwise or logical operations => and or xor and_not
	fmt.Printf("%v\n", a&b)
	fmt.Printf("%v\n", a|b)
	fmt.Printf("%v\n", a^b)
	fmt.Printf("%v\n", a&^b)

	// Left shift and Right shift
	a = 8
	fmt.Printf("%v\n", a<<3) // shift 3 places to left so 2^3 * 2^3 = 8*8 = 64
	b = 8
	fmt.Printf("%v\n", a>>3) // shift 3 places to right so 2^3 / 2^3 = 8/8 = 1

	// Floats => float32(-+ 1.18E-38 to -+3.34E38) and float64(-+2.24E-308 to -+1.18E308)

	x := 12.3
	y := 10.5
	fmt.Printf("%v\n", x+y)
	fmt.Printf("%v\n", x-y)
	fmt.Printf("%v\n", x*y)
	fmt.Printf("%v\n", x/y)

	var c complex64 = 1 + 2i
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("real = %v, image=%v\n", real(c), imag(c))

	var d complex64 = 2 + 3i
	fmt.Printf("%v\n", c+d)
	fmt.Printf("%v\n", c-d)
	fmt.Printf("%v\n", c*d)
	fmt.Printf("%v\n", c/d)

	var e complex64 = complex(4, 12)
	fmt.Printf("%v, %T, real = %v, image=%v\n", e, e, real(e), imag(e))

	// text type
	s := "Hello Ish"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s) // here s[2] return character ascii value
	fmt.Printf("%v, %T\n", string(s[2]), s[2]) // here string(s[2]) return character

	s1 := " Good morning"
	fmt.Printf("%v\n", s+s1) // string concatenation

	slice := []byte(s) // converting in to slice of bytes
	fmt.Printf("%v, %T\n", slice, slice) // string concatenation

	var r rune = 'm'  // rune use single quote
	fmt.Printf("%v, %T\n", r, r)

}
