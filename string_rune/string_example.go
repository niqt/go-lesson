package stringrune

import (
	"fmt"
	"unicode/utf8"
)

func string_base() {
	// String semplice ASCII
	s1 := "hello"
	fmt.Printf("String: %s\n", s1)
	fmt.Printf("Length in bytes: %d\n", len(s1))               // 5 bytes
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(s1)) // 5 runes

	// String con caratteri non-ASCII
	s2 := "ciao"
	fmt.Printf("\nString: %s\n", s2)
	fmt.Printf("Length in bytes: %d\n", len(s2))               // 4 bytes
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(s2)) // 4 runes

	// String con emoji e caratteri speciali
	s3 := "Hello 世界 🌍"
	fmt.Printf("\nString: %s\n", s3)
	fmt.Printf("Length in bytes: %d\n", len(s3))               // 16 bytes!
	fmt.Printf("Rune count: %d\n", utf8.RuneCountInString(s3)) // 9 runes

	fmt.Println("\n--- Breakdown of s3 ---")
	// 'H' = 1 byte
	// 'e' = 1 byte
	// 'l' = 1 byte
	// 'l' = 1 byte
	// 'o' = 1 byte
	// ' ' = 1 byte
	// '世' = 3 bytes
	// '界' = 3 bytes
	// ' ' = 1 byte
	// '🌍' = 4 bytes
	// Total: 16 bytes
}

func rune_base() {
	// Rune literal (single character in single quotes)
	var r1 rune = 'A'
	var r2 rune = '世'
	var r3 rune = '🌍'

	fmt.Printf("r1: %c (Unicode: U+%04X, decimal: %d)\n", r1, r1, r1)
	fmt.Printf("r2: %c (Unicode: U+%04X, decimal: %d)\n", r2, r2, r2)
	fmt.Printf("r3: %c (Unicode: U+%04X, decimal: %d)\n", r3, r3, r3)

	// Output:
	// r1: A (Unicode: U+0041, decimal: 65)
	// r2: 世 (Unicode: U+4E16, decimal: 19990)
	// r3: 🌍 (Unicode: U+1F30D, decimal: 127757)
}

func conversion() {
	s := "Ciao 世界"

	// Convert string to slice of runes
	runes := []rune(s)
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Runes: %v\n", runes)
	fmt.Printf("Number of runes: %d\n", len(runes)) // 7

	// Now we can safely access by character index
	fmt.Printf("First character: %c\n", runes[0]) // C
	fmt.Printf("Last character: %c\n", runes[6])  // 界

	// Modify runes
	runes[0] = 'K'
	runes[5] = '世'

	// Convert back to string
	newString := string(runes)
	fmt.Printf("Modified string: %s\n", newString) // Kiao 世世
}
