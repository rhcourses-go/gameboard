package board

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// TestMakeBoard prüft die korrekte Verwendung von MakeBoard.
func TestMakeBoard(t *testing.T) {
	b1 := MakeBoard("*", 1, 1)
	b1_expected := Board{MakeRow("*", 1)}

	b2 := MakeBoard("*", 2, 2)
	b2_expected := Board{
		MakeRow("*", 2),
		MakeRow("*", 2),
	}

	b3 := MakeBoard("*", 3, 3)
	b3_expected := Board{
		MakeRow("*", 3),
		MakeRow("*", 3),
		MakeRow("*", 3),
	}

	b4 := MakeBoard("*", 2, 4)
	b4_expected := Board{
		MakeRow("*", 4),
		MakeRow("*", 4),
	}

	b5 := MakeBoard("*", 4, 2)
	b5_expected := Board{
		MakeRow("*", 2),
		MakeRow("*", 2),
		MakeRow("*", 2),
		MakeRow("*", 2),
	}

	errormsgtemplate := "MakeBoard returned unexpected result.\nexpected rows:\n%v\ngot:\n%v."
	boardrows := func(b Board) string {
		rows := make([]string, b.Rows())
		for i, row := range b {
			rows[i] = fmt.Sprintf("  %#v", row)
		}
		return strings.Join(rows, "\n")
	}

	if !reflect.DeepEqual(b1, b1_expected) {
		t.Errorf(errormsgtemplate, boardrows(b1_expected), boardrows(b1))
	}
	if !reflect.DeepEqual(b2, b2_expected) {
		t.Errorf(errormsgtemplate, boardrows(b2_expected), boardrows(b2))
	}
	if !reflect.DeepEqual(b3, b3_expected) {
		t.Errorf(errormsgtemplate, boardrows(b3_expected), boardrows(b3))
	}
	if !reflect.DeepEqual(b4, b4_expected) {
		t.Errorf(errormsgtemplate, boardrows(b4_expected), boardrows(b4))
	}
	if !reflect.DeepEqual(b5, b5_expected) {
		t.Errorf(errormsgtemplate, boardrows(b5_expected), boardrows(b5))
	}
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

// TestMakeBoardFromStrings prüft die korrekte Verwendung von MakeBoardFromStrings.
func TestMakeBoardFromStrings(t *testing.T) {
	b1 := MakeBoardFromStrings("ABC", "DEF", "GHI")

	expectedRows := "  ABC\n  DEF\n  GHI\n"
	actualRows := ""
	for _, row := range b1 {
		actualRows += fmt.Sprintf("  %s\n", row.AsString())
	}

	if !reflect.DeepEqual(actualRows, expectedRows) {
		t.Errorf(
			"MakeBoardFromStrings returned unexpected result.\nexpected rows:\n%s\ngot:\n%s.",
			expectedRows,
			actualRows,
		)
	}
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

// ExampleBoard_GetRow zeigt die korrekte Verwendung von GetRow.
func ExampleBoard_GetRow() {
	b := MakeBoardFromStrings(
		"ABC",
		"DEF",
		"GHI",
	)
	fmt.Println(b.GetRow(0).AsString())
	fmt.Println(b.GetRow(1).AsString())
	fmt.Println(b.GetRow(2).AsString())

	// Output:
	// ABC
	// DEF
	// GHI
}

// ExampleBoard_GetCol zeigt die korrekte Verwendung von GetCol.
func ExampleBoard_GetCol() {
	b := MakeBoardFromStrings(
		"ABC",
		"DEF",
		"GHI",
	)
	fmt.Println(b.GetCol(0).AsString())
	fmt.Println(b.GetCol(1).AsString())
	fmt.Println(b.GetCol(2).AsString())

	// Output:
	// ADG
	// BEH
	// CFI
}

// ExampleBoard_GetDiagRight zeigt die korrekte Verwendung von GetDiagRight.
func ExampleBoard_GetDiagRight() {
	b := MakeBoardFromStrings(
		"ABCDE",
		"FGHIJ",
		"KLMNO",
	)
	fmt.Println(b.GetDiagRight(-2).AsString())
	fmt.Println(b.GetDiagRight(-1).AsString())
	fmt.Println(b.GetDiagRight(0).AsString())
	fmt.Println(b.GetDiagRight(1).AsString())
	fmt.Println(b.GetDiagRight(2).AsString())
	fmt.Println(b.GetDiagRight(3).AsString())
	fmt.Println(b.GetDiagRight(4).AsString())

	// Output:
	// K
	// FL
	// AGM
	// BHN
	// CIO
	// DJ
	// E

}

// ExampleBoard_GetDiagLeft zeigt die korrekte Verwendung von GetDiagLeft.
func ExampleBoard_GetDiagLeft() {
	b := MakeBoardFromStrings(
		"ABCDE",
		"FGHIJ",
		"KLMNO",
	)
	fmt.Println(b.GetDiagLeft(0).AsString())
	fmt.Println(b.GetDiagLeft(1).AsString())
	fmt.Println(b.GetDiagLeft(2).AsString())
	fmt.Println(b.GetDiagLeft(3).AsString())
	fmt.Println(b.GetDiagLeft(4).AsString())
	fmt.Println(b.GetDiagLeft(5).AsString())
	fmt.Println(b.GetDiagLeft(6).AsString())

	// Output:
	// A
	// BF
	// CGK
	// DHL
	// EIM
	// JN
	// O
}
