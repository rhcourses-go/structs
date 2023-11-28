package sports

// Match ist ein Datentyp, der Metadaten zu einer Sportveranstaltung speichern soll.
type Match struct {
	Home     string // Name der Heimmannschaft
	Visitors string // Name der Auswärtsmannschaft

	Location string // Name des Austragungsorts (Stadt, Stadion o.Ä.)
	Result   int    // 0: Unentschieden, 1: Heimmannschaft gewinnt, 2: Gäste gewinnen
}

// SetHomeTeam erwartet einen String und schreibt ihn in das Attribut Home.
func (m *Match) SetHomeTeam(name string) {
	m.Home = name
}

// SetVisitorTeam erwartet einen String und schreibt ihn in das Attribut Visitors.
func (m *Match) SetVisitorTeam(name string) {
	m.Visitors = name
}

// SetLocation erwartet einen String und schreibt ihn in das Attribut Location.
// Außerdem setzt sie auch die Heimmanschaft auf den Namen des Austragungsorts,
// wenn das Attribut Home noch leer ist.
func (m *Match) SetLocation(name string) {
	m.Location = name
	if m.Home == "" {
		m.Home = name
	}
}

/* Anmerkung:
Die Methoden hier haben Pointer-Receiver.
D.h. wenn sie aufgerufen werden, liegt m als Pointer auf ein Match vor.
Dies bedeutet, dass sich Änderungen als *Seiteneffekt* auf das Objekt auswirken,
auf dem die Methode aufgerufen wurde.

Um zu sehen, was das bedeutet, entfernen Sie die Pointer (die Sterne) von den
Receiver-Variablen. Sie werden dann in den Tests sehen, dass die Aufrufe der Methoden
dann keinen Effekt mehr haben.
Dies liegt daran, dass die Match-Objekte dann beim Aufruf der Methoden
kopiert werden.
*/
