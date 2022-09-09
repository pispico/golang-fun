package main

import "fmt"

//Examples of operations with slices

func main() {
	slice := []int{3, 4, 5}
	fmt.Printf("Original slice: %d\n", slice)

	//apending in slice
	a := []int{6, 7}
	slice = append(slice, a...)
	fmt.Printf("Slice after appending: %d\n", slice)

	//prepending in slice
	p := []int{0, 1, 2}
	slice = append(p, slice...)
	fmt.Printf("Slice after prepending: %d\n", slice)

	//Removing element from slice
	idx := 4
	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Printf("Slice after index 4 removed: %d\n", slice)

	//In this example every 2 sequencial elements is a group
	// lets return the 3rd group, in this case would be 5 and 6
	myslice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	offset := 2 * (3 - 1) //2 is the number of members in each group, 3 is the group number
	fmt.Printf("The 3rd group is%d\n", myslice[offset:offset+2])

	//sum one in slice element
	addoneslice := []int{0, 0, 0}
	addoneslice[1]++
	fmt.Printf("After sum +1 in element 1 %d\n", addoneslice)

	//Changing value of slice element 2
	addoneslice[2] = 100
	fmt.Printf("After changing the slice element 2 %d\n", addoneslice)

	//Reversing array
	arrayReverse := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Array before reverse %d\n", arrayReverse)
	reverse(arrayReverse[:])
	fmt.Printf("Array after reverse %d\n", arrayReverse)

	//Rotating left array 5 elements
	arrayRotate := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Array before rotate left 5 elements %d\n", arrayRotate)
	reverse(arrayRotate[:5])
	reverse(arrayRotate[5:])
	reverse(arrayRotate[:])
	fmt.Printf("Array after rotate left 5 elements %d\n", arrayRotate)

}

//This function rotates an array
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
