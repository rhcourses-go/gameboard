package board

import "strings"

// RowContainsN erwartet eine Zeile row, einen String s und eine Anzahl n.
// Liefert true, wenn der String in der Zeile row n mal hintereinander vorkommt.
// Liefert false, wenn keine solche Teilsequenz existiert
// oder wenn die Parameterangaben ungültig sind.
// Löst eine Panic aus, wenn row ungültig ist.
func (b Board) RowContainsN(row int, s string, n int) bool {
	PanicIfPositionInvalid(row, b.Rows())
	return b.GetRow(row).ContainsN(s, n)
}

// ColContainsN erwartet eine Spalte col, einen String s und eine Anzahl n.
// Liefert true, wenn der String in der Spalte col n mal hintereinander vorkommt.
// Liefert false, wenn keine solche Teilsequenz existiert
// oder wenn die Parameterangaben ungültig sind.
// Löst eine Panic aus, wenn col ungültig ist.
func (b Board) ColContainsN(col int, s string, n int) bool {
	PanicIfPositionInvalid(col, b.Cols())
	return b.GetCol(col).ContainsN(s, n)
}

// DiagRightContainsN erwartet Spalte col, einen String s und eine Anzahl n.
// Liefert true, wenn der String in der Diagonalen von links oben nach rechts unten
// beginnend in Spalte col n mal hintereinander vorkommt.
// Liefert false, wenn keine solche Teilsequenz existiert
// oder wenn die Parameterangaben ungültig sind.
func (b Board) DiagRightContainsN(col int, s string, n int) bool {
	return b.GetDiagRight(col).ContainsN(s, n)
}

// DiagLeftContainsN erwartet Spalte col, einen String s und eine Anzahl n.
// Liefert true, wenn der String in der Diagonalen von rechts oben nach links unten
// beginnend in Spalte col n mal hintereinander vorkommt.
// Liefert false, wenn keine solche Teilsequenz existiert
// oder wenn die Parameterangaben ungültig sind.
func (b Board) DiagLeftContainsN(col int, s string, n int) bool {
	return b.GetDiagLeft(col).ContainsN(s, n)
}

// RowContainsString erwartet eine Zeile row und einen String s.
// Liefert true, wenn s als Teilsequenz in der Zeile vorkommt.
// D.h. wenn die Zeile alle Zeichen aus s hintereinander und
// in der richtigen Reihenfolge enthält.
// Liefert false, wenn keine solche Teilsequenz existiert.
func (b Board) RowContainsString(row int, s string) bool {
	return b.GetRow(row).ContainsString(s)
}

// ColContainsString erwartet eine Spalte col und einen String s.
// Liefert true, wenn s als Teilsequenz in der Spalte vorkommt.
// D.h. wenn die Spalte alle Zeichen aus s hintereinander und
// in der richtigen Reihenfolge enthält.
// Liefert false, wenn keine solche Teilsequenz existiert.
func (b Board) ColContainsString(col int, s string) bool {
	return b.GetCol(col).ContainsString(s)
}

// DiagRightContainsString erwartet Spalte col und einen String s.
// Liefert true, wenn s als Teilsequenz in der Diagonalen von links oben nach rechts unten
// beginnend in Spalte col vorkommt.
// D.h. wenn die Diagonale alle Zeichen aus s hintereinander und
// in der richtigen Reihenfolge enthält.
// Liefert false, wenn keine solche Teilsequenz existiert.
func (b Board) DiagRightContainsString(col int, s string) bool {
	return b.GetDiagRight(col).ContainsString(s)
}

// DiagLeftContainsString erwartet Spalte col und einen String s.
// Liefert true, wenn s als Teilsequenz in der Diagonalen von rechts oben nach links unten
// beginnend in Spalte col vorkommt.
// D.h. wenn die Diagonale alle Zeichen aus s hintereinander und
// in der richtigen Reihenfolge enthält.
// Liefert false, wenn keine solche Teilsequenz existiert.
func (b Board) DiagLeftContainsString(col int, s string) bool {
	return b.GetDiagLeft(col).ContainsString(s)
}

// RowContainsOnly erwartet eine Zeile row und ein Zeichen c.
// Liefert true, wenn die Zeile nur aus dem Zeichen c besteht.
// Löst eine Panik aus, wenn der String oder die Zeilenangabe ungültig ist.
func (b Board) RowContainsOnly(row int, c string) bool {
	PanicIfStringInvalid(c)

	sequence := strings.Repeat(c, b.Cols())
	return b.RowContainsString(row, sequence)
}

// ColContainsOnly erwartet eine Spalte col und ein Zeichen c.
// Liefert true, wenn die Spalte nur aus dem Zeichen c besteht.
// Löst eine Panik aus, wenn der String oder die Spaltenangabe ungültig ist.
func (b Board) ColContainsOnly(col int, c string) bool {
	PanicIfStringInvalid(c)

	sequence := strings.Repeat(c, b.Rows())
	return b.ColContainsString(col, sequence)
}

// DiagRightContainsOnly erwartet eine Spalte col und ein Zeichen c.
// Liefert true, wenn die Diagonale von links oben nach rechts unten
// beginnend in Spalte col nur aus dem Zeichen c besteht.
// Löst eine Panik aus, wenn der String oder die Spaltenangabe ungültig ist.
func (b Board) DiagRightContainsOnly(col int, c string) bool {
	PanicIfStringInvalid(c)

	diagRight := b.GetDiagRight(col)
	for _, cell := range diagRight {
		if cell != c {
			return false
		}
	}
	return true
}

// DiagLeftContainsOnly erwartet eine Spalte col und ein Zeichen c.
// Liefert true, wenn die Diagonale von rechts oben nach links unten
// beginnend in Spalte col nur aus dem Zeichen c besteht.
// Löst eine Panik aus, wenn der String oder die Spaltenangabe ungültig ist.
func (b Board) DiagLeftContainsOnly(col int, c string) bool {
	PanicIfStringInvalid(c)

	diagLeft := b.GetDiagLeft(col)
	for _, cell := range diagLeft {
		if cell != c {
			return false
		}
	}
	return true
}
