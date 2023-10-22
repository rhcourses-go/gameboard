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

// TestRow_String prüft die korrekte Verwendung von Row.String.
func TestRow_String(t *testing.T) {
	r1 := MakeRowFromString("XOXX")
	r1_expected := MakeRow(" ", 4)
	r1_expected[0] = "X"
	r1_expected[1] = "O"
	r1_expected[2] = "X"
	r1_expected[3] = "X"

	if !reflect.DeepEqual(r1, r1_expected) {
		t.Errorf("MakeRowFromString returned unexpected result.\n  expected: %#v\n  got:      %#v", r1_expected, r1)
	}
}

// TestRow_String_panic_with_empty_string testet,
// ob Row.String bei leerem String eine Panic auslöst.
func TestRow_String_panic_with_empty_string(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Row.String should have panicked.")
		}
	}()
	MakeRowFromString("")
}

// Anmerkung: Tests für ungültige Zeichen mit MakeRowFromString sind derzeit nicht
// sinnvoll, da die Funktion nur einzelne Zeichen erwartet und alle einzelnen Zeichen
// erlaubt sind. Sollte sich diese Anforderung an die Zeichen einmal ändern, müssen
// entsprechende Tests hinzugefügt werden.

// TestRow_AsString prüft die korrekte Verwendung von Row.AsString.
func TestRow_AsString(t *testing.T) {
	r1 := MakeRowFromString("XOXX")
	r1_expected := "XOXX"

	r2 := MakeRowFromString("X  X")
	r2_expected := "X  X"

	if r1.AsString() != r1_expected {
		t.Errorf("Row.AsString returned unexpected result.\n  expected: %v\n  got:      %v", r1_expected, r1.AsString())
	}

	if r2.AsString() != r2_expected {
		t.Errorf("Row.AsString returned unexpected result.\n  expected: %v\n  got:      %v", r2_expected, r2.AsString())
	}
}
