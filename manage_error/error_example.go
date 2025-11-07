package manageerror

import (
	"errors"
	"fmt"
	"strconv"
)

// Function that returns error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function with formatted error
func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("invalid age: %d (must be >= 0)", age)
	}
	if age > 150 {
		return fmt.Errorf("invalid age: %d (must be <= 150)", age)
	}
	return nil
}

// Custom error type
type ValidationError struct {
	Field string
	Value interface{}
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s (value: %v)",
		e.Field, e.Msg, e.Value)
}

func validateUser(name string, age int) error {
	if name == "" {
		return &ValidationError{
			Field: "name",
			Value: name,
			Msg:   "cannot be empty",
		}
	}
	if age < 18 {
		return &ValidationError{
			Field: "age",
			Value: age,
			Msg:   "must be at least 18",
		}
	}
	return nil
}

func error_example() {
	// Basic error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("10 / 2 = %.2f\n", result)

	// Error case
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	}

	// If with error evaluation (VERY COMMON IDIOM!)
	if result, err := divide(20, 4); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("20 / 4 = %.2f\n", result)
	}

	// Multiple operations with early return
	if err := validateAge(25); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Age 25 is valid")

	if err := validateAge(-5); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Custom error type
	fmt.Println("\n--- Custom errors ---")
	if err := validateUser("", 20); err != nil {
		fmt.Printf("Error: %v\n", err)

		// Type assertion to get custom error details
		if validErr, ok := err.(*ValidationError); ok {
			fmt.Printf("Field: %s, Value: %v\n", validErr.Field, validErr.Value)
		}
	}

	// Wrapping errors (Go 1.13+)
	fmt.Println("\n--- Wrapped errors ---")
	if num, err := strconv.Atoi("abc"); err != nil {
		wrappedErr := fmt.Errorf("failed to parse number: %w", err)
		fmt.Printf("Wrapped error: %v\n", wrappedErr)

		// Unwrap to check original error
		if errors.Is(wrappedErr, strconv.ErrSyntax) {
			fmt.Println("Original error was a syntax error")
		}
	} else {
		fmt.Printf("Parsed: %d\n", num)
	}

	// Panic and recover (used rarely!)
	fmt.Println("\n--- Panic and recover ---")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	// This would crash without recover
	panic("something went terribly wrong!")
}
