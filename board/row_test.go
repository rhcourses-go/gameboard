package board

import (
	"reflect"
	"testing"
)

// TestMakeRow prüft die korrekte Verwendung von MakeRow.
func TestMakeRow(t *testing.T) {
	r1 := MakeRow(" ", 3)
	r1_expected := Row{" ", " ", " "}

	r2 := MakeRow("*", 4)
	r2_expected := Row{"*", "*", "*", "*"}

	if !reflect.DeepEqual(r1, r1_expected) {
		t.Errorf("MakeRow returned unexpected result.\n  expected: %#v\n  got: %#v.", r1_expected, r1)
	}
	if !reflect.DeepEqual(r2, r2_expected) {
		t.Errorf("MakeRow returned unexpected result.\n  expected: %#v\n  got:      %#v", r2_expected, r2)
	}
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
