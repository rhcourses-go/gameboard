package board

import (
	"fmt"
	"testing"
)

// ExampleMakeRow zeigt die korrekte Verwendung von MakeRow.
func ExampleMakeRow() {
	fmt.Println(MakeRow(" ", 3))
	fmt.Println(MakeRow("*", 4))

	// Output:
	// |   |   |   |
	// | * | * | * | * |
}

// TestMakeRowInvalidLength testet, ob MakeRow eine Panic auslöst,
// wenn die Länge 0 oder negativ ist.
func TestMakeRow_PanicWithZeroLength(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeRow should have panicked.")
		}
	}()
	MakeRow("A", 0)
}

// TestMakeRow_panic_with_negative_length testet,
// ob MakeRow bei negativer Länge eine Panic auslöst.
func TestMakeRow_panic_with_negative_length(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeRow should have panicked.")
		}
	}()
	MakeRow("A", -1)
}

// TestMakeRow_panic_with_invalid_fill testet,
// ob MakeRow bei ungültigem Füllzeichen eine Panic auslöst.
func TestMakeRow_panic_with_invalid_fill(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeRow should have panicked.")
		}
	}()
	MakeRow("XX", 1)
}
