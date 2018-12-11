package main

import (
	"fmt"
)

func main() {

	array := []int{1, 5, 7, 2, 4, 6, 3}
	fmt.Println("Input array is ", array)
	length := len(array)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println("Sorted array is ", array)
}
