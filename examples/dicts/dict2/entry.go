package dict2

import "fmt"

/*
Entry ist ein Container für Einträge in einem Deutsch-Englisch-Wörterbuch.
In dieser ersten einfachen Form wird in jedem Eintrag immer genau ein
deutsches Wort (De) auf genau ein englisches Wort (En) abgebildet.
*/
type Entry struct {
	De string // Das deutsche Wort in diesem Eintrag
	En string // Das englische Wort in diesem Eintrag
}

/*
NewEntry ist ein Konstruktor für Entry-Objekte.
Diese Funktion erwartet die Werte für De und En und liefert einen Entry mit diesen Werten.
*/
func NewEntry(de, en string) Entry {
	return Entry{de, en}
}

/*
String ist eine Methode, die eine menschenlesbare Repräsenation des Eintrags liefert.
Ist eine solche Methode für einen Typ vorhanden, so wird diese z.B. von fmt.Println
verwendet.
Die Methode hat keine normalen Parameter, dafür aber einen Receiver.
Dieser gibt an, zu welchem Typ die Methode gehört.
Ansonsten ist er wie ein normaler Parameter.
*/
func (entry Entry) String() string {
	if entry.De == "" || entry.En == "" {
		return "[unvollständiger Eintrag]"
	}

	return fmt.Sprintf(
		"Deutsch: %s, Englisch: %s",
		entry.De,
		entry.En,
	)
}
