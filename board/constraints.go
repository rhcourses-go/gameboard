package board

import "fmt"

// StringValid erwartet einen String und liefert true, wenn dieser
// ein gültiges Zeichen für ein Spielfeld ist.
// Gültige Zeichen sind alle einelementigen Strings.
func StringValid(s string) bool {
	return len(s) == 1
}

// PanicIfStringInvalid erwartet einen String und löst eine Panic aus,
// wenn dieser kein gültiges Zeichen für ein Spielfeld ist.
func PanicIfStringInvalid(s string) {
	if !StringValid(s) {
		panic("Der String " + s + " ist kein gültiges Zeichen für ein Spielfeld.")
	}
}

// LengthValid erwartet eine Länge und liefert true, wenn diese für ein Spielfeld gültig ist.
// Gültige Längen sind alle positiven Zahlen.
func LengthValid(length int) bool {
	return length > 0
}

// PanicIfLengthInvalid erwartet eine Länge und löst eine Panic aus,
// wenn diese für ein Spielfeld ungültig ist.
func PanicIfLengthInvalid(length int) {
	if !LengthValid(length) {
		panic("Die Länge einer Zeile muss positiv sein.")
	}
}

// PanicIfPositionInvalid erwartet eine Position und eine obere Grenze.
// Löst eine Panic aus, wenn die Position ungültig ist.
func PanicIfPositionInvalid(position, upper int) {
	if position < 0 || position >= upper {
		panic(fmt.Sprintf("Die Position %d ist ungültig.", position))
	}
}
