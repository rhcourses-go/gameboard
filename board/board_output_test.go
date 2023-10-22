package board

import "fmt"

// ExampleBoard_String zeigt die korrekte Verwendung von Board.String.
func ExampleBoard_String() {
	fmt.Println(MakeBoard(" ", 1, 1))
	fmt.Println(MakeBoard(" ", 2, 2))
	fmt.Println(MakeBoard(" ", 3, 3))
	fmt.Println(MakeBoard(" ", 2, 4))
	fmt.Println(MakeBoard(" ", 4, 2))

	// Output:
	// +---+
	// |   |
	// +---+
	//
	// +---+---+
	// |   |   |
	// +---+---+
	// |   |   |
	// +---+---+
	//
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	// |   |   |   |
	// +---+---+---+
	//
	// +---+---+---+---+
	// |   |   |   |   |
	// +---+---+---+---+
	// |   |   |   |   |
	// +---+---+---+---+
	//
	// +---+---+
	// |   |   |
	// +---+---+
	// |   |   |
	// +---+---+
	// |   |   |
	// +---+---+
	// |   |   |
	// +---+---+
}
