package board

import "fmt"

// ExampleBoard_RowContainsN zeigt die korrekte Verwendung von Board.RowContainsN.
func ExampleBoard_RowContainsN() {
	b := MakeBoardFromStrings(
		"XOXXXXOOO",
		"XXXXXXXXX",
		"OOOOOOOOO",
	)

	fmt.Println(b.RowContainsN(0, "X", 1))
	fmt.Println(b.RowContainsN(0, "X", 2))
	fmt.Println(b.RowContainsN(0, "X", 4))
	fmt.Println(b.RowContainsN(0, "X", 5))

	fmt.Println(b.RowContainsN(1, "O", 3))
	fmt.Println(b.RowContainsN(1, "X", 5))

	fmt.Println(b.RowContainsN(2, "O", 3))
	fmt.Println(b.RowContainsN(2, "X", 5))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// true
	// true
	// false
}

// ExampleBoard_ColContainsN zeigt die korrekte Verwendung von Board.ColContainsN.
func ExampleBoard_ColContainsN() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"OXX",
	)

	fmt.Println(b.ColContainsN(0, "X", 0))
	fmt.Println(b.ColContainsN(0, "X", 1))
	fmt.Println(b.ColContainsN(0, "X", 2))
	fmt.Println(b.ColContainsN(0, "X", 3))
	fmt.Println(b.ColContainsN(0, "X", 4))

	fmt.Println()

	fmt.Println(b.ColContainsN(1, "X", 0))
	fmt.Println(b.ColContainsN(1, "X", 1))
	fmt.Println(b.ColContainsN(1, "X", 2))
	fmt.Println(b.ColContainsN(1, "X", 3))
	fmt.Println(b.ColContainsN(1, "X", 4))

	fmt.Println()

	fmt.Println(b.ColContainsN(2, "X", 0))
	fmt.Println(b.ColContainsN(2, "X", 1))
	fmt.Println(b.ColContainsN(2, "X", 2))
	fmt.Println(b.ColContainsN(2, "X", 3))
	fmt.Println(b.ColContainsN(2, "X", 4))

	// Output:
	// true
	// true
	// false
	// false
	// false
	//
	// true
	// true
	// true
	// false
	// false
	//
	// true
	// true
	// false
	// false
	// false
}

// ExampleBoard_DiagRightContainsN zeigt die korrekte Verwendung von Board.DiagRightContainsN.
func ExampleBoard_DiagRightContainsN() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"OXX",
	)

	fmt.Println(b.DiagRightContainsN(0, "X", 0))
	fmt.Println(b.DiagRightContainsN(0, "X", 1))
	fmt.Println(b.DiagRightContainsN(0, "X", 2))
	fmt.Println(b.DiagRightContainsN(0, "X", 3))
	fmt.Println(b.DiagRightContainsN(0, "X", 4))

	fmt.Println()

	fmt.Println(b.DiagRightContainsN(1, "O", 0))
	fmt.Println(b.DiagRightContainsN(1, "O", 1))
	fmt.Println(b.DiagRightContainsN(1, "O", 2))
	fmt.Println(b.DiagRightContainsN(1, "O", 3))

	// Output:
	// true
	// true
	// true
	// true
	// false
	//
	// true
	// true
	// true
	// false
}

// ExampleBoard_DiagLeftContainsN zeigt die korrekte Verwendung von Board.DiagLeftContainsN.
func ExampleBoard_DiagLeftContainsN() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"OXX",
	)

	fmt.Println(b.DiagLeftContainsN(0, "X", 0))
	fmt.Println(b.DiagLeftContainsN(0, "X", 1))
	fmt.Println(b.DiagLeftContainsN(0, "X", 2))

	fmt.Println()

	fmt.Println(b.DiagLeftContainsN(1, "O", 0))
	fmt.Println(b.DiagLeftContainsN(1, "O", 1))
	fmt.Println(b.DiagLeftContainsN(1, "O", 2))
	fmt.Println(b.DiagLeftContainsN(1, "O", 3))

	// Output:
	// true
	// true
	// false
	//
	// true
	// true
	// true
	// false
}

// ExampleBoard_RowContainsString zeigt die korrekte Verwendung von Board.RowContainsString.
func ExampleBoard_RowContainsString() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"OXX",
	)

	fmt.Println(b.RowContainsString(0, "XOX"))
	fmt.Println(b.RowContainsString(1, "OXO"))
	fmt.Println(b.RowContainsString(2, "OXX"))
	fmt.Println(b.RowContainsString(2, "XOX"))

	// Output:
	// true
	// true
	// true
	// false
}

// ExampleBoard_ColContainsString zeigt die korrekte Verwendung von Board.ColContainsString.
func ExampleBoard_ColContainsString() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"OXX",
	)

	fmt.Println(b.ColContainsString(0, "XOO"))
	fmt.Println(b.ColContainsString(1, "XOX"))
	fmt.Println(b.ColContainsString(1, "OXX"))
	fmt.Println(b.ColContainsString(2, "OXX"))
	fmt.Println(b.ColContainsString(2, "XOX"))

	// Output:
	// true
	// false
	// true
	// false
	// true
}

// ExampleBoard_DiagRightContainsString zeigt die korrekte Verwendung von Board.DiagRightContainsString.
func ExampleBoard_DiagRightContainsString() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"OXX",
	)

	fmt.Println(b.DiagRightContainsString(0, "X"))
	fmt.Println(b.DiagRightContainsString(0, "XX"))
	fmt.Println(b.DiagRightContainsString(0, "XXX"))
	fmt.Println(b.DiagRightContainsString(0, "XO"))

	fmt.Println()

	fmt.Println(b.DiagRightContainsString(1, "O"))
	fmt.Println(b.DiagRightContainsString(1, "OO"))
	fmt.Println(b.DiagRightContainsString(1, "OOO"))
	fmt.Println(b.DiagRightContainsString(1, "OOX"))

	fmt.Println()

	fmt.Println(b.DiagRightContainsString(2, "X"))
	fmt.Println(b.DiagRightContainsString(2, "O"))
	fmt.Println(b.DiagRightContainsString(2, "XX"))

	// Output:
	// true
	// true
	// true
	// false
	//
	// true
	// true
	// false
	// false
	//
	// true
	// false
	// false

}

// ExampleBoard_DiagLeftContainsString zeigt die korrekte Verwendung von Board.DiagLeftContainsString.
func ExampleBoard_DiagLeftContainsString() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"OXX",
	)

	fmt.Println(b.DiagLeftContainsString(0, "X"))
	fmt.Println(b.DiagLeftContainsString(0, "XX"))

	fmt.Println()

	fmt.Println(b.DiagLeftContainsString(1, "O"))
	fmt.Println(b.DiagLeftContainsString(1, "OO"))
	fmt.Println(b.DiagLeftContainsString(1, "OOO"))
	fmt.Println(b.DiagLeftContainsString(1, "OOX"))

	fmt.Println()

	fmt.Println(b.DiagLeftContainsString(2, "X"))
	fmt.Println(b.DiagLeftContainsString(2, "O"))
	fmt.Println(b.DiagLeftContainsString(2, "XX"))
	fmt.Println(b.DiagLeftContainsString(2, "XXO"))
	fmt.Println(b.DiagLeftContainsString(2, "XXX"))

	// Output:
	// true
	// false
	//
	// true
	// true
	// false
	// false
	//
	// true
	// true
	// true
	// true
	// false
}

// ExampleBoard_RowContainsOnly zeigt die korrekte Verwendung von Board.RowContainsOnly.
func ExampleBoard_RowContainsOnly() {
	b := MakeBoardFromStrings(
		"XOX",
		"OOO",
		"XXX",
	)

	fmt.Println(b.RowContainsOnly(0, "X"))
	fmt.Println(b.RowContainsOnly(0, "O"))
	fmt.Println()
	fmt.Println(b.RowContainsOnly(1, "X"))
	fmt.Println(b.RowContainsOnly(1, "O"))
	fmt.Println()
	fmt.Println(b.RowContainsOnly(2, "X"))
	fmt.Println(b.RowContainsOnly(2, "O"))

	// Output:
	// false
	// false
	//
	// false
	// true
	//
	// true
	// false
}

// ExampleBoard_ColContainsOnly zeigt die korrekte Verwendung von Board.ColContainsOnly.
func ExampleBoard_ColContainsOnly() {
	b := MakeBoardFromStrings(
		"XOX",
		"OOX",
		"XOX",
	)

	fmt.Println(b.ColContainsOnly(0, "X"))
	fmt.Println(b.ColContainsOnly(0, "O"))
	fmt.Println()
	fmt.Println(b.ColContainsOnly(1, "X"))
	fmt.Println(b.ColContainsOnly(1, "O"))
	fmt.Println()
	fmt.Println(b.ColContainsOnly(2, "X"))
	fmt.Println(b.ColContainsOnly(2, "O"))

	// Output:
	// false
	// false
	//
	// false
	// true
	//
	// true
	// false
}

// ExampleBoard_DiagRightContainsOnly zeigt die korrekte Verwendung von Board.DiagRightContainsOnly.
func ExampleBoard_DiagRightContainsOnly() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"OXX",
	)

	fmt.Println(b.DiagRightContainsOnly(0, "X"))
	fmt.Println(b.DiagRightContainsOnly(0, "O"))
	fmt.Println()
	fmt.Println(b.DiagRightContainsOnly(1, "X"))
	fmt.Println(b.DiagRightContainsOnly(1, "O"))
	fmt.Println()
	fmt.Println(b.DiagRightContainsOnly(2, "X"))
	fmt.Println(b.DiagRightContainsOnly(2, "O"))
	fmt.Println()
	fmt.Println(b.DiagRightContainsOnly(3, "X"))
	fmt.Println(b.DiagRightContainsOnly(3, "O"))
	fmt.Println()
	fmt.Println(b.DiagRightContainsOnly(-1, "X"))
	fmt.Println(b.DiagRightContainsOnly(-1, "O"))

	// Output:
	// true
	// false
	//
	// false
	// true
	//
	// true
	// false
	//
	// true
	// true
	//
	// false
	// false

}

// ExampleBoard_DiagLeftContainsOnly zeigt die korrekte Verwendung von Board.DiagLeftContainsOnly.
func ExampleBoard_DiagLeftContainsOnly() {
	b := MakeBoardFromStrings(
		"XOX",
		"OXO",
		"XXO",
	)

	fmt.Println(b.DiagLeftContainsOnly(0, "X"))
	fmt.Println(b.DiagLeftContainsOnly(0, "O"))
	fmt.Println()
	fmt.Println(b.DiagLeftContainsOnly(1, "X"))
	fmt.Println(b.DiagLeftContainsOnly(1, "O"))
	fmt.Println()
	fmt.Println(b.DiagLeftContainsOnly(2, "X"))
	fmt.Println(b.DiagLeftContainsOnly(2, "O"))
	fmt.Println()
	fmt.Println(b.DiagLeftContainsOnly(3, "X"))
	fmt.Println(b.DiagLeftContainsOnly(3, "O"))
	fmt.Println()
	fmt.Println(b.DiagLeftContainsOnly(-1, "X"))
	fmt.Println(b.DiagLeftContainsOnly(-1, "O"))

	// Output:
	// true
	// false
	//
	// false
	// true
	//
	// true
	// false
	//
	// false
	// false
	//
	// true
	// true
}
