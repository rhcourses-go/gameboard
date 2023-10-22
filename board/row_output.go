package board

import (
	"strings"
)

// String() liefert eine Zeichenkette, die die Zeile repräsentiert.
// Die Zeichenkette enthält Trennzeichen, um die einzelnen Einträge
// der Zeile voneinander zu trennen.
func (r Row) String() string {
	result := strings.Join(r, " | ")
	return "| " + result + " |"
}
