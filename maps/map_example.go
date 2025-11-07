package maps

import "fmt"

func map_example() {
	// Declaration and initialization with make
	m1 := make(map[string]int)
	m1["apple"] = 5
	m1["banana"] = 3
	fmt.Printf("m1: %v\n", m1)

	// Literal initialization
	m2 := map[string]int{
		"red":   1,
		"green": 2,
		"blue":  3,
	}
	fmt.Printf("m2: %v\n", m2)

	// Access elements
	fmt.Printf("m2[\"red\"] = %d\n", m2["red"])

	// Check if key exists (IMPORTANT!)
	value, exists := m2["yellow"]
	if exists {
		fmt.Printf("yellow exists: %d\n", value)
	} else {
		fmt.Printf("yellow doesn't exist (value: %d - zero value)\n", value)
	}

	// Accessing non-existent key returns zero value
	fmt.Printf("m2[\"purple\"] = %d (zero value)\n", m2["purple"])

	// Modify value
	m2["red"] = 10
	fmt.Printf("After modification: %v\n", m2)

	// Delete key
	delete(m2, "green")
	fmt.Printf("After delete: %v\n", m2)

	// Iterate over map (ORDER IS NOT GUARANTEED!)
	fmt.Println("\n--- Iterating over map ---")
	for key, value := range m2 {
		fmt.Printf("%s: %d\n", key, value)
	}

	// Length of map
	fmt.Printf("\nLength: %d\n", len(m2))

	// Map of structs
	type Person struct {
		Name string
		Age  int
	}

	people := map[string]Person{
		"nicola": {Name: "Nicola", Age: 30},
		"mario":  {Name: "Mario", Age: 25},
	}
	fmt.Printf("\nPeople: %v\n", people)

	// Map with pointer values (to modify structs in place)
	peoplePtr := map[string]*Person{
		"nicola": {Name: "Nicola", Age: 30},
	}
	peoplePtr["nicola"].Age = 31 // Can modify directly
	fmt.Printf("Modified age: %v\n", peoplePtr["nicola"])

	// Nested maps
	nested := map[string]map[string]int{
		"fruits": {
			"apple":  5,
			"banana": 3,
		},
		"vegetables": {
			"carrot": 10,
			"potato": 7,
		},
	}
	fmt.Printf("\nNested map: %v\n", nested)
	fmt.Printf("fruits[apple]: %d\n", nested["fruits"]["apple"])
}
