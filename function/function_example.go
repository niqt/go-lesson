package function

import "fmt"

// Simple function
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with return value
func add(a, b int) int {
	return a + b
}

// Multiple parameters of same type (shorthand)
func multiply(x, y, z int) int {
	return x * y * z
}

// Multiple return values (VERY common in Go!)
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Named return values
func swap(x, y string) (first, second string) {
	first = y
	second = x
	return // "naked return" - returns named values
}

// Variadic function (variable number of arguments)
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Function as value
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Anonymous function (closure)
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Pass by VALUE - doesn't modify original
func incrementValue(x int) {
	x = x + 10
	fmt.Printf("Inside incrementValue: %d\n", x)
}

// Pass by POINTER - modifies original
func incrementPointer(x *int) {
	*x = *x + 10
	fmt.Printf("Inside incrementPointer: %d\n", *x)
}

// Structs by value - doesn't modify original
type Person struct {
	Name string
	Age  int
}

func birthdayValue(p Person) {
	p.Age++
	fmt.Printf("Inside birthdayValue: %s is %d\n", p.Name, p.Age)
}

// Structs by pointer - modifies original
func birthdayPointer(p *Person) {
	p.Age++ // Go automatically dereferences, no need for (*p).Age
	fmt.Printf("Inside birthdayPointer: %s is %d\n", p.Name, p.Age)
}

// Slice behaves like reference (shares underlying array)
func modifySlice(s []int) {
	s[0] = 999 // Modifies underlying array
	fmt.Printf("Inside modifySlice: %v\n", s)
}

// Map behaves like reference
func modifyMap(m map[string]int) {
	m["x"] = 100 // Modifies original map
	fmt.Printf("Inside modifyMap: %v\n", m)
}
