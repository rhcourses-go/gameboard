package board

import "fmt"

// ExampleRowStr zeigt, wie Zeilen auf der Konsole ausgegeben werden.
func ExampleRow_String() {
	r1 := MakeRow(" ", 3)
	fmt.Println(r1.String())
	fmt.Println(r1)

	r2 := MakeRow("*", 4)
	fmt.Println(r2.String())
	fmt.Println(r2)

	// Output:
	// |   |   |   |
	// |   |   |   |
	// | * | * | * | * |
	// | * | * | * | * |
}
