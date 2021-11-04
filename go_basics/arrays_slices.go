package main

import (
	"fmt"
)

func main() {
	marks := []int{30, 40, 50} // array declaration and initialization
	fmt.Printf("%v\n", marks)

	newMarks := [...]int{10, 20, 30, 40} // ... will automatically takes the size of array
	fmt.Printf("%v\n", newMarks)

	var snames [3]string // declaring an empty array
	fmt.Printf("%v\n", snames)

	snames[0] = "Raju"
	snames[1] = "Shaam"
	snames[2] = "Babu Rao"

	fmt.Printf("%v\n", snames)
	fmt.Printf("Ye %v ka style hai\n", snames[2])
	fmt.Printf("%v\n", len(snames))

	// var arr1 [3]int = [3]int{1,2,3}
	arr1 := [...]int{1, 2, 3}
	arr2 := arr1 // arr2 holds the copy of array unlike references in python
	arr2[0] = 4
	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", arr2)

	arr3 := &arr1 // here arr3 holds the address of arr1 working on pointer-reference concept
	arr3[0] = 5   // this change will also reflect in arr1
	fmt.Printf("%v\n", arr1)
	fmt.Println(arr3)

	slice1 := []int{1, 2, 3} // slices, they are by default reference types
	fmt.Println(slice1)

	fmt.Println(len(slice1))
	fmt.Printf("Capacity = %v\n", cap(slice1))

	slice2 := slice1
	fmt.Println(slice1)
	fmt.Println("slice2 = ", slice2)
	slice2[0] = 5
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice4 := slice3[:]   // [:] copy all elements of slice3 in slice 4
	slice5 := slice3[3:]  // [3:] copy all elements of slice3 in slice 5 from index 4
	slice6 := slice3[:6]  // [:6] copy all elements of slice3 in slice 6 upto index 5
	slice7 := slice3[3:6] // [3:6] copy all elements of slice3 in slice 7 from index 2 to index 5
	slice3[3] = 12
	fmt.Println("slice3 = ", slice3)
	fmt.Println("slice4 = ", slice4)
	fmt.Println("slice5 = ", slice5)
	fmt.Println("slice6 = ", slice6)
	fmt.Println("slice7 = ", slice7)

	s := make([]int, 3)  // creating slice using make function
	fmt.Println(s)
	fmt.Println("Length = ", len(s))
	fmt.Println("Capacity = ", cap(s))

	// s := make([]int, 3, 100)  // creating slice of 3 elements from an array of 100 elements

	s = []int{}
	s = append(s, 1,2,3,4,5)
	fmt.Println(s)
	fmt.Println("Length = ", len(s))
	fmt.Println("Capacity = ", cap(s))

	s = append(s, 11,12,13,14,15)
	fmt.Println(s)
	fmt.Println("Length = ", len(s))
	fmt.Println("Capacity = ", cap(s))

	s = append(s, 21,22,23,24,25,26)
	fmt.Println(s)
	fmt.Println("Length = ", len(s))
	fmt.Println("Capacity = ", cap(s))


	s = append(s, []int{2,3,4,5}...)
	fmt.Println(s)
	fmt.Println("Length = ", len(s))
	fmt.Println("Capacity = ", cap(s))

/*
	// 2D array
	var matrix2d [2][2]int = [2][2]int{
		[2]int{0, 1}, [2]int{2, 3},
	} // declaration and initialization in a single line
	fmt.Printf("%v\n", matrix2d)
	// initialization in seprate line
	matrix2d[0] = [2]int{1, 2}
	matrix2d[1] = [2]int{3, 4}
	fmt.Printf("%v\n", matrix2d)
*/
}
