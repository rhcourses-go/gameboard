package board

import (
	"strings"
)

// ContainsN erwartet einen String und eine Zahl n.
// Liefert true, wenn der String in der Zeile n mal hintereinander vorkommt.
// Liefert false, wenn keine solche Teilsequenz existiert
// oder wenn die Parameterangaben ungültig sind.
func (r Row) ContainsN(s string, n int) bool {
	if !StringValid(s) || n < 0 || n > len(r) {
		return false
	}

	rowAsString := r.AsString()
	substring := strings.Repeat(s, n)
	return strings.Contains(rowAsString, substring)
}

// ContainsString erwartet einen String s.
// Liefert true, wenn s als Teilsequenz in der Zeile vorkommt.
// D.h. wenn die Zeile alle Zeichen aus s hintereinander und
// in der richtigen Reihenfolge enthält.
// Liefert false, wenn keine solche Teilsequenz existiert.
func (r Row) ContainsString(s string) bool {
	rowString := r.AsString()
	return strings.Contains(rowString, s)
}
