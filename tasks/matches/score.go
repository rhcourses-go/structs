package matches

import "fmt"

// Score ist ein Datentyp, der einen Spielstand eines Spiels speichern soll.
// Gemeint ist etwas wie "1:0" oder "3:2".
type Score struct {
	Home     int
	Visitors int
}

// NewScore erzeugt ein neues Score-Objekt mit den übergebenen Werten.
func NewScore(home, visitors int) Score {
	return Score{Home: home, Visitors: visitors}
}

// String gibt den Score als String in der Form "1:0" zurück.
func (s Score) String() string {
	return fmt.Sprintf("%d:%d", s.Home, s.Visitors)
}

// Result gibt zurück, wer das Spiel gewonnen hat.
// 0: Unentschieden, 1: Heimmannschaft gewinnt, 2: Gäste gewinnen
func (s Score) Result() int {
	// begin:hint
	// Prüfen Sie, ob die Heimmannschaft mehr Tore geschossen hat als die Gäste.
	// Wenn ja, geben Sie 1 zurück.
	// Wenn nein, prüfen Sie, ob die Gäste mehr Tore geschossen haben als die Heimmannschaft.
	// Wenn ja, geben Sie 2 zurück.
	// Wenn nein, geben Sie 0 zurück.
	// end:hint
	// begin:solution
	if s.Home > s.Visitors {
		return 1
	} else if s.Home < s.Visitors {
		return 2
	} else {
		return 0
	}
	// end:solution
	// iftask: return 0
}
