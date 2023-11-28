package dict1

/*
Entry ist ein Container für Einträge in einem Deutsch-Englisch-Wörterbuch.
In dieser ersten einfachen Form wird in jedem Eintrag immer genau ein
deutsches Wort (De) auf genau ein englisches Wort (En) abgebildet.
*/
type Entry struct {
	De string // Das deutsche Wort in diesem Eintrag
	En string // Das englische Wort in diesem Eintrag
}
