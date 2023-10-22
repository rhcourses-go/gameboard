package board

/*
Board ist ein Typ, der ein rechteckiges Spielfeld repräsentiert.
Das Spielfeld besteht aus Zeilen, die wiederum aus Einträgen bestehen.
*/
type Board []Row

// MakeBoard erwartet einen String 'fill' und zwei Längen.
// Erzeugt und liefert ein neues Spielfeld mit der angegebenen
// Anzahl an Zeilen und Einträgen pro Zeile.
// Alle Einträge des Spielfelds sind mit dem String 'fill' gefüllt.
// Wenn die Längen nicht positiv oder 'fill' nicht erlaubt ist, wird eine Panic ausgelöst.
func MakeBoard(fill string, rows, cols int) Board {
	PanicIfLengthInvalid(rows)
	PanicIfLengthInvalid(cols)
	PanicIfStringInvalid(fill)

	board := make(Board, rows)
	for i := range board {
		board[i] = MakeRow(fill, cols)
	}
	return board
}

// MakeBoardFromStrings erwartet eine Slice aus Strings 'rows'
// und erzeugt und liefert ein neues Spielfeld mit den Einträgen aus 'rows'.
// Wenn einer der Einträge nicht erlaubt ist, wird eine Panic ausgelöst.
func MakeBoardFromStrings(rows ...string) Board {
	PanicIfLengthInvalid(len(rows))
	board := make(Board, len(rows))
	for i, row := range rows {
		board[i] = MakeRowFromString(row)
	}
	return board
}

// Rows liefert die Anzahl der Zeilen.
func (b Board) Rows() int {
	return len(b)
}

// Cols liefert die Anzahl der Einträge pro Zeile.
func (b Board) Cols() int {
	return len(b[0])
}

// IsSquare liefert true, wenn das Spielfeld quadratisch ist.
func (b Board) IsSquare() bool {
	return b.Rows() == b.Cols()
}
