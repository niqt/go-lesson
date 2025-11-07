package controlflow

import "fmt"

func if_example() {
	// Simple if
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// If with else
	if x > 15 {
		fmt.Println("x is greater than 15")
	} else {
		fmt.Println("x is not greater than 15")
	}

	// If-else if-else chain
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// If with initialization statement (VERY COMMON!)
	if num := 42; num > 40 {
		fmt.Printf("num (%d) is greater than 40\n", num)
	} // num is only in scope inside the if block
}

func for_example() {
	nums := []int{10, 20, 30, 40, 50}

	// Classic for with index
	fmt.Println("--- Classic for loop ---")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, nums[i])
	}

	// For-range (index and value)
	fmt.Println("\n--- For-range (index + value) ---")
	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// For-range (only value)
	fmt.Println("\n--- For-range (only value) ---")
	for _, v := range nums {
		fmt.Printf("Value: %d\n", v)
	}

	// For-range (only index)
	fmt.Println("\n--- For-range (only index) ---")
	for i := range nums {
		fmt.Printf("Index: %d\n", i)
	}

	// For-range with string (iterates by RUNE, not byte!)
	fmt.Println("\n--- For-range on string ---")
	s := "Go 世界"
	for i, r := range s {
		fmt.Printf("Byte pos: %d, Rune: %c (U+%04X)\n", i, r, r)
	}
}
