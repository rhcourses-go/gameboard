package board

import (
	"fmt"
	"testing"
)

// ExampleMakeBoard zeigt die korrekte Verwendung von MakeBoard.
func ExampleMakeBoard() {
	fmt.Println(MakeBoard("*", 1, 1))
	fmt.Println(MakeBoard("*", 2, 2))
	fmt.Println(MakeBoard("*", 3, 3))
	fmt.Println(MakeBoard("*", 2, 4))

	// Output:
	// +---+
	// | * |
	// +---+
	//
	// +---+---+
	// | * | * |
	// +---+---+
	// | * | * |
	// +---+---+
	//
	// +---+---+---+
	// | * | * | * |
	// +---+---+---+
	// | * | * | * |
	// +---+---+---+
	// | * | * | * |
	// +---+---+---+
	//
	// +---+---+---+---+
	// | * | * | * | * |
	// +---+---+---+---+
	// | * | * | * | * |
	// +---+---+---+---+
}

// TestMakeBoardZeroColumns testet, ob MakeBoard eine Panic auslöst,
// wenn die Anzahl der Zeilen 0 ist.
func TestMakeBoardZeroRows(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeBoard should have panicked.")
		}
	}()
	MakeBoard("A", 0, 1)
}

// TestMakeBoardZeroColumns testet, ob MakeBoard eine Panic auslöst,
// wenn die Anzahl der Spalten 0 ist.
func TestMakeBoardZeroColumns(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeBoard should have panicked.")
		}
	}()
	MakeBoard("A", 1, 0)
}

// TestMakeBoardNegativeRows testet, ob MakeBoard eine Panic auslöst,
// wenn die Anzahl der Zeilen negativ ist.
func TestMakeBoardNegativeRows(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeBoard should have panicked.")
		}
	}()
	MakeBoard("A", -1, 1)
}

// TestMakeBoardNegativeColumns testet, ob MakeBoard eine Panic auslöst,
// wenn die Anzahl der Spalten negativ ist.
func TestMakeBoardNegativeColumns(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeBoard should have panicked.")
		}
	}()
	MakeBoard("A", 1, -1)
}

// TestMakeBoardInvalidFill testet, ob MakeBoard eine Panic auslöst,
// wenn das Füllzeichen ungültig ist.
func TestMakeBoardInvalidFill(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MakeBoard should have panicked.")
		}
	}()
	MakeBoard("XX", 1, 1)
}

// ExampleBoard_Rows zeigt die korrekte Verwendung von Rows.
func ExampleBoard_Rows() {
	fmt.Println(MakeBoard("*", 1, 1).Rows())
	fmt.Println(MakeBoard("*", 2, 2).Rows())
	fmt.Println(MakeBoard("*", 3, 3).Rows())
	fmt.Println(MakeBoard("*", 2, 4).Rows())

	// Output:
	// 1
	// 2
	// 3
	// 2
}

// ExampleBoard_Cols zeigt die korrekte Verwendung von Cols.
func ExampleBoard_Cols() {
	fmt.Println(MakeBoard("*", 1, 1).Cols())
	fmt.Println(MakeBoard("*", 2, 2).Cols())
	fmt.Println(MakeBoard("*", 3, 3).Cols())
	fmt.Println(MakeBoard("*", 2, 4).Cols())

	// Output:
	// 1
	// 2
	// 3
	// 4
}

// ExampleBoard_IsSquare zeigt die korrekte Verwendung von IsSquare.
func ExampleBoard_IsSquare() {
	fmt.Println(MakeBoard("*", 1, 1).IsSquare())
	fmt.Println(MakeBoard("*", 2, 2).IsSquare())
	fmt.Println(MakeBoard("*", 3, 3).IsSquare())
	fmt.Println(MakeBoard("*", 2, 4).IsSquare())
	fmt.Println(MakeBoard("*", 4, 2).IsSquare())

	// Output:
	// true
	// true
	// true
	// false
	// false
}
