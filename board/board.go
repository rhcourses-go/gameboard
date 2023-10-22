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

// GetRow erwartet einen Index und liefert die Zeile mit dem angegebenen Index.
// Wenn der Index ungültig ist, wird eine Panic ausgelöst.
func (b Board) GetRow(index int) Row {
	return b[index]
}

// GetCol erwartet einen Index und liefert die Spalte mit dem angegebenen Index.
// Wenn der Index ungültig ist, wird eine Panic ausgelöst.
func (b Board) GetCol(index int) Row {
	col := make(Row, b.Rows())
	for i := range col {
		col[i] = b[i][index]
	}
	return col
}

// GetDiagRight erwartet eine Spalte col und liefert die Diagonale,
// die in der Spalte col beginnt und nach rechts unten verläuft.
// Löst keine Panic aus, wenn die Spalte ungültig ist, sondern liefert
// ggf. eine Teil-Diagonale, die in einer späteren Zeile beginnt.
// Diese Teil-Diagonale kann auch leer sein.
func (b Board) GetDiagRight(col int) Row {
	rows, cols := b.Rows(), b.Cols()
	diag := make(Row, 0)
	for i := 0; i < rows && i+col < cols; i++ {
		if i+col >= 0 {
			diag = append(diag, b[i][i+col])
		}
	}
	return diag
}

// GetDiagLeft erwartet eine Spalte col und liefert die Diagonale,
// die in der Spalte col beginnt und nach links unten verläuft.
// Löst keine Panic aus, wenn die Spalte ungültig ist, sondern liefert
// ggf. eine Teil-Diagonale, die in einer späteren Zeile beginnt.
// Diese Teil-Diagonale kann auch leer sein.
func (b Board) GetDiagLeft(col int) Row {
	diag := make(Row, 0)
	for i := 0; i < b.Rows() && col-i >= 0; i++ {
		if col-i < b.Cols() {
			diag = append(diag, b[i][col-i])
		}
	}
	return diag
}
