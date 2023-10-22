package board

import (
	"fmt"
	"strings"
)

// String gibt das Spielfeld als String zurück.
// Dabei werden die einzelnen Zeichen der Zeilen durch einen Trenner getrennt
// und die Zeilen haben ebenfalls Trennzeilen.
func (b Board) String() string {
	rowSeparator := strings.Repeat("+---", b.Cols()) + "+\n"
	rowStrings := make([]string, b.Rows())

	for i, row := range b {
		rowStrings[i] = fmt.Sprintf("%s\n", row.String())
	}

	return fmt.Sprintf(
		"%s%s%s",
		rowSeparator,
		strings.Join(rowStrings, rowSeparator),
		rowSeparator,
	)
}
