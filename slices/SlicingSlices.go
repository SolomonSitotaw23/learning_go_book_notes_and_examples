package slices

import "fmt"

// this is example of how to create slices from slices
func SlicingSlices() {
	slice := []int{1, 2, 3, 4, 5}

	slice1 := slice[1:4]

	slice2 := slice[0:2]

	slice3 := slice[:2]

	slice4 := slice[3:]

	fmt.Println(slice)
	fmt.Println("[1:4]", slice1)
	fmt.Println("[0:2]", slice2)
	fmt.Println("[:2]", slice3)
	fmt.Println("[3:]", slice4)

}
