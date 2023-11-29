package matches

import "fmt"

// Match ist ein Datentyp, der das Ergebnis eines Sportspiels speichern soll.
type Match struct {
	home     string
	visitors string

	score Score
}

// NewMatch erzeugt ein neues Match-Objekt mit den übergebenen Werten.
func NewMatch(home, visitors string, score Score) Match {
	return Match{home: home, visitors: visitors, score: score}
}

// HomeName gibt den Namen der Heimmannschaft zurück.
// Falls die Heimmannschaft gewonnen hat, wird der Name in Sternchen gesetzt.
func (m Match) HomeName() string {
	// begin:hint
	// Verwenden Sie die Methode Result() des Score-Typs, um zu prüfen,
	// ob die Heimmannschaft gewonnen hat.
	// end:hint
	// begin:solution
	if m.score.Result() == 1 {
		return fmt.Sprintf("*%s*", m.home)
	}
	// end:solution
	return m.home
}

// VisitorName gibt den Namen der Auswärtsmannschaft zurück.
// Falls die Auswärtsmannschaft gewonnen hat, wird der Name in Sternchen gesetzt.
func (m Match) VisitorName() string {
	// begin:hint
	// Gehen Sie analog zu HomeName vor.
	// end:hint
	// begin:solution
	if m.score.Result() == 2 {
		return fmt.Sprintf("*%s*", m.visitors)
	}
	// end:solution
	return m.visitors
}

// String gibt das Match als String in der Form "FC Freiburg - *Borussia Bremen*: 1:2" zurück.
// Dabei soll immer der Name der Heimmannschaft zuerst stehen und der Name des
// Gewinners in Sternchen gesetzt werden.
func (m Match) String() string {
	// begin:hint
	// Verwenden Sie die Methoden HomeName und VisitorName, um die Namen der
	// Mannschaften zu erhalten.
	// Verwenden Sie das Feld m.score, um den Spielstand zu erhalten.
	// end:hint
	// begin:solution
	return fmt.Sprintf("%s - %s: %s", m.HomeName(), m.VisitorName(), m.score)
	// end:solution
	// iftask: return ""
}

// Winner gibt den Namen des Gewinners zurück.
// Wenn es keinen Gewinner gibt, wird "unentschieden" zurückgegeben.
func (m Match) Winner() string {
	// begin:hint
	// Verwenden Sie m.score.Result(), um zu prüfen, ob es einen Gewinner gibt.
	// end:hint
	// begin:solution
	switch m.score.Result() {
	case 1:
		return m.home
	case 2:
		return m.visitors
	default:
		return "unentschieden"
	}
	// end:solution
	// iftask: return ""
}
