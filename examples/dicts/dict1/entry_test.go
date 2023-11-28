package dict1

import "fmt"

/*
ExampleEntry_definition demonstriert, wie man Objekte vom Typ Entry definiert.

Variablen mit selbstdefinierten Typen können grundsätzlich auf die gleiche Weise
definiert werden, wie die eingebauten Typen auch.

D.h. man kann die Schreibweise var <NAME> <TYP> verwenden, oder man gibt den
Variablennamen gefolgt von `:=` und einer Initialisierung ein.

Da es sich um ein Struct handelt, muss man beim Initialisieren die Werte für die Felder
angeben. Dies geschieht, wie auch bei Arrays und Slices, mit geschweiften Klammern.
Dabei ist es auch möglich, die Namen der Felder explizit mit anzugeben.

Lässt man die Klammern leer, werden die Felder Null-Initialisiert, also mit ihrem
jeweiligen Standardwert.
*/
func ExampleEntry_definition() {

	// Definition eines default-initialisierten Entry-Objekts mit expliziter Typangabe:
	var entry1 Entry

	// Definition eines default-initialisierten Entry-Objekts mit Typ-Inferenz:
	entry2 := Entry{}

	// Definition eines Entry-Objekts mit Typ-Inferenz und Angabe der Werte:
	entry3 := Entry{"Haus", "house"}

	// Definition eines Entry-Objekts mit Typ-Inferenz und Angabe der Werte:
	entry4 := Entry{En: "bicycle", De: "Fahrrad"}

	// Ausgabe der erzeugten Entry-Objekte
	// (in Go-Code-Syntax, damit die leeren Strings sichtbar sind):
	fmt.Printf("entry1: %#v\n", entry1)
	fmt.Printf("entry2: %#v\n", entry2)
	fmt.Printf("entry3: %#v\n", entry3)
	fmt.Printf("entry4: %#v\n", entry4)

	// Output:
	// entry1: dict1.Entry{De:"", En:""}
	// entry2: dict1.Entry{De:"", En:""}
	// entry3: dict1.Entry{De:"Haus", En:"house"}
	// entry4: dict1.Entry{De:"Fahrrad", En:"bicycle"}
}

/*
ExampleEntry_usage demonstriert, wie man Objekte vom Typ Entry verwendet.

Auf die Felder eines Eintrags kann man mit der Punkt-Notation zugreifen.
Hier wird zunächst ein Eintrag erzeugt, seine Felder ausgegeben,
verändert und wieder ausgegeben.
*/
func ExampleEntry_usage() {

	// Definition eines Eintrags mit nicht ganz korrekter englischer Übersetzung.
	entry1 := Entry{De: "Pflanze", En: "flower"}

	// Ausgabe des Eintrags.
	// Zuerst als ganzes (%+v gibt ein Struct mit Feld-Namen aus)
	// und danach die beiden Felder einzeln:

	fmt.Println("entry1 direkt nach der Definition:")
	fmt.Printf("entry1: %+v\n", entry1)
	fmt.Printf("entry1.De: %v\n", entry1.De)
	fmt.Printf("entry1.En: %v\n", entry1.En)
	fmt.Println()

	// Wir korrigieren die Übersetzung von "Pflanze"
	// und geben den Eintrag anschließend wieder aus:
	entry1.En = "plant"
	fmt.Println("entry1 nach der Änderung von entry1.En:")
	fmt.Printf("entry1: %+v\n", entry1)
	fmt.Printf("entry1.De: %v\n", entry1.De)
	fmt.Printf("entry1.En: %v\n", entry1.En)

	// Output:
	// entry1 direkt nach der Definition:
	// entry1: {De:Pflanze En:flower}
	// entry1.De: Pflanze
	// entry1.En: flower
	//
	// entry1 nach der Änderung von entry1.En:
	// entry1: {De:Pflanze En:plant}
	// entry1.De: Pflanze
	// entry1.En: plant
}
