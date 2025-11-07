package array

import "fmt"

func array_example() {
	// Declaration with size
	var arr1 [5]int
	fmt.Printf("arr1: %v (zero values)\n", arr1) // [0 0 0 0 0]

	// Declaration with initialization
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr2: %v\n", arr2) // [1 2 3 4 5]

	// Partial initialization (rest are zero values)
	arr3 := [5]int{1, 2}
	fmt.Printf("arr3: %v\n", arr3) // [1 2 0 0 0]

	// Let compiler count the size
	arr4 := [...]int{10, 20, 30}
	fmt.Printf("arr4: %v (length: %d)\n", arr4, len(arr4)) // [10 20 30] (length: 3)

	// Access elements
	arr2[0] = 100
	fmt.Printf("Modified arr2: %v\n", arr2) // [100 2 3 4 5]

	// Length and capacity are the same
	fmt.Printf("len(arr2): %d, cap(arr2): %d\n", len(arr2), cap(arr2)) // 5, 5
}
