package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Input array is ", array)
	start, end := 0, len(array)-1
	key := 2
	for start <= end {

		if array[start] == key {
			fmt.Println("Found at", array[start])
			return
		}
		if array[end] == key {
			fmt.Println("Found at", array[end])
			return
		}
		middle := (start + end) / 2
		if array[middle] == key {
			fmt.Println("Found at", array[middle])
			return
		}
		if array[middle] > key {
			end = middle - 1
		} else {
			start = middle + 1
		}

	}
}
