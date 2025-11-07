package slice

import "fmt"

func slice_base1() {
	// Create slice with make
	s1 := make([]int, 5) // length 5, capacity 5
	fmt.Printf("s1: %v (len: %d, cap: %d)\n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 10) // length 3, capacity 10
	fmt.Printf("s2: %v (len: %d, cap: %d)\n", s2, len(s2), cap(s2))

	// Create slice with literal
	s3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s3: %v (len: %d, cap: %d)\n", s3, len(s3), cap(s3))

	// Empty slice
	var s4 []int
	fmt.Printf("s4: %v (len: %d, cap: %d, nil: %t)\n", s4, len(s4), cap(s4), s4 == nil)
}

func slice_base2() {
	// A slice is a struct with 3 fields:
	// - pointer to underlying array
	// - length (number of elements)
	// - capacity (size of underlying array from the slice's start)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Create slice from array
	s := arr[2:7] // Elements from index 2 to 6 (7 is excluded)

	fmt.Printf("Original array: %v\n", arr)
	fmt.Printf("Slice s: %v (len: %d, cap: %d)\n", s, len(s), cap(s))
	// s = [2 3 4 5 6], len=5, cap=8 (from index 2 to end of array)

	// Modifying slice modifies the underlying array
	s[0] = 999
	fmt.Printf("After modifying s[0]:\n")
	fmt.Printf("Array: %v\n", arr) // [0 1 999 3 4 5 6 7 8 9]
	fmt.Printf("Slice: %v\n", s)   // [999 3 4 5 6]
}
