package board

import (
	"fmt"
	"testing"
)

// ExampleStringValid zeigt die korrekte Verwendung von StringValid.
func ExampleStringValid() {
	fmt.Println(StringValid("X"))
	fmt.Println(StringValid("O"))
	fmt.Println(StringValid(" "))
	fmt.Println(StringValid("XX"))
	fmt.Println(StringValid(""))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

// TestPanikIfStringInvalid testet, ob PanicIfStringInvalid eine Panic auslöst,
// wenn der String ungültig ist.
func TestPanicIfStringInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PanicIfStringInvalid should have panicked.")
		}
	}()
	PanicIfStringInvalid("XX")
}

// ExampleLengthValid zeigt die korrekte Verwendung von LengthValid.
func ExampleLengthValid() {
	fmt.Println(LengthValid(1))
	fmt.Println(LengthValid(2))
	fmt.Println(LengthValid(0))
	fmt.Println(LengthValid(-1))

	// Output:
	// true
	// true
	// false
	// false
}

// TestPanicIfLengthZero testet, ob TestPanicIfLengthZero eine Panic auslöst,
// wenn die Länge ungültig ist.
func TestPanicIfLengthZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PanicIfLengthInvalid should have panicked.")
		}
	}()
	PanicIfLengthInvalid(0)
}

// TestPanicIfLengthNegative testet, ob TestPanicIfLengthNegative eine Panic auslöst,
// wenn die Länge ungültig ist.
func TestPanicIfLengthNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PanicIfLengthInvalid should have panicked.")
		}
	}()
	PanicIfLengthInvalid(-1)
}

// TestPanicIfPositionInvalid_negative testet, ob PanicIfPositionInvalid eine Panic auslöst,
// wenn die Position negativ ist.
func TestPanicIfPositionInvalid_negative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PanicIfPositionInvalid should have panicked.")
		}
	}()
	PanicIfPositionInvalid(-1, 1)
}

// TestPanicIfPositionInvalid_too_large testet, ob PanicIfPositionInvalid eine Panic auslöst,
// wenn die Position zu groß ist.
func TestPanicIfPositionInvalid_too_large(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PanicIfPositionInvalid should have panicked.")
		}
	}()
	PanicIfPositionInvalid(1, 1)
}
