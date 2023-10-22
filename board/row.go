package board

/*
Row bezeichnet eine Zeile eines Spielfelds: eine Slice aus Strings.

Anmerkung: Der Typ Row ist ein Alias für []string und damit eigentlich
etwas allgemeineres als eine Zeile eines Spielfelds.
Wir verwenden diesen Typ z.B. auch für unvollständige Zeilen oder
wenn wir einzelne Spalten oder Diagonalen eines Spielfelds betrachten.

Die Grundidee ist, dass jeder Eintrag in der Slice genau ein Zeichen
des Spielfelds ist, z.B. "X", "O" oder " " (Leerzeichen).
Wir verwenden dennoch den Typ string, da die Einträge so einfacher zu handhaben und auszugeben sind.
Einschränkungen der Einträge werden durch die Methoden des Typs Row sichergestellt.
*/
type Row []string

// MakeRow erwartet einen String 'fill' und eine Länge.
// Erzeugt und liefert eine neue Zeile mit der angegebenen Länge.
// Alle Einträge der Zeile sind mit dem String 'fill' gefüllt.
// Wenn die Länge nicht positiv oder 'fill' nicht erlaubt ist, wird eine Panic ausgelöst.
func MakeRow(fill string, length int) Row {
	PanicIfLengthInvalid(length)
	PanicIfStringInvalid(fill)

	row := make(Row, length)
	for i := range row {
		row[i] = fill
	}
	return row
}
