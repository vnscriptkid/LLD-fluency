package main

import "fmt"

// Arrays are passed by value
func modifyArray(arr [3]int) {
	arr[0] = 100
}

// Slices are passed by reference
func modifySlice(slice []int) {
	slice[0] = 100
}

func main() {
	println("Hello, playground")
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	slice := []int{1, 2, 3}

	modifyArray(arr)
	modifySlice(slice)

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("slice: %v\n", slice)
}
