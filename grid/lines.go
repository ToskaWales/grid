package grid

// Diese Aufgaben liefern eine Linie aus dem Spielfeld: eine Zeile,
// eine Spalte oder eine Diagonale. Jede Linie ist eine neue Liste;
// wird sie geändert, bleibt das Spielfeld unverändert.

// Row erwartet ein Spielfeld und die Nummer `index` einer Zeile
// und liefert die Zellen dieser Zeile von links nach rechts.
//
// TODO
func Row(cells [][]string, index int) []string {
	result := []string{}

	for _, value := range cells[index] {
		result = append(result, value)
	}
	return result
}

// Col erwartet ein Spielfeld und die Nummer `index` einer Spalte
// und liefert die Zellen dieser Spalte von oben nach unten.
//
// TODO
func Col(cells [][]string, index int) []string {
	return nil
}

// Diagonal erwartet ein Spielfeld und liefert die Diagonale von links oben
// nach rechts unten: die Zellen, deren Zeile und Spalte gleich sind.
// Bei einem Spielfeld, das nicht quadratisch ist, endet sie am Rand.
//
// TODO
func Diagonal(cells [][]string) []string {
	return nil
}

// AntiDiagonal erwartet ein Spielfeld und liefert die Diagonale von rechts oben
// nach links unten. Bei einem Spielfeld, das nicht quadratisch ist,
// endet sie am Rand.
//
// TODO
func AntiDiagonal(cells [][]string) []string {
	return nil
}
