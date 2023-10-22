package board

import (
	"fmt"
)

// ExampleRow_ContainsN zeigt die korrekte Verwendung von Row.ContainsSequence.
func ExampleRow_ContainsN() {
	r1 := MakeRowFromString("XOXXXXOOO")

	fmt.Println(r1.ContainsN("X", 1))
	fmt.Println(r1.ContainsN("X", 2))
	fmt.Println(r1.ContainsN("X", 4))
	fmt.Println(r1.ContainsN("X", 5))

	fmt.Println(r1.ContainsN("O", 3))
	fmt.Println(r1.ContainsN("O", 4))

	fmt.Println(r1.ContainsN("X", 0))

	fmt.Println(r1.ContainsN("XX", 1))

	// Output:
	// true
	// true
	// true
	// false
	// true
	// false
	// true
	// false
}

// ExampleRow_ContainsString zeigt die korrekte Verwendung von Row.ContainsString.
func ExampleRow_ContainsString() {
	r1 := MakeRowFromString("XOXXXXOOO")

	// Positive Fälle
	fmt.Println(r1.ContainsString("X"))
	fmt.Println(r1.ContainsString("O"))
	fmt.Println(r1.ContainsString("XO"))
	fmt.Println(r1.ContainsString("OX"))
	fmt.Println(r1.ContainsString("XX"))
	fmt.Println(r1.ContainsString("OO"))
	fmt.Println(r1.ContainsString("XOX"))

	fmt.Println()

	// Negative Fälle
	fmt.Println(r1.ContainsString("OXO"))
	fmt.Println(r1.ContainsString("XOOOO"))

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
	//
	// false
	// false
}
