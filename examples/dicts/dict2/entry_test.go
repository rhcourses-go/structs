package dict2

import "fmt"

/*
ExampleEntry_definition demonstriert, wie man Objekte vom Typ Entry definiert.

Zusätzlich zu den Beispielen aus dict1 wird hier noch der neue Konstruktor NewEntry
verwendet.

Außerdem ist die Ausgabe verändert:
Dadurch, dass eine Methode String existiert, wird diese auch bei der Ausgabe verwendet.
In den fmt.Printf-aufrufen verwenden wir nun das normale "%v".
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

	// Definition eines Entry-Objekts mittels NewEntry:
	entry5 := NewEntry("Auto", "car")

	// Ausgabe der erzeugten Entry-Objekte
	// (in Go-Code-Syntax, damit die leeren Strings sichtbar sind):
	fmt.Printf("entry1: %v\n", entry1)
	fmt.Printf("entry2: %v\n", entry2)
	fmt.Printf("entry3: %v\n", entry3)
	fmt.Printf("entry4: %v\n", entry4)
	fmt.Printf("entry5: %v\n", entry5)

	// Output:
	// entry1: [unvollständiger Eintrag]
	// entry2: [unvollständiger Eintrag]
	// entry3: Deutsch: Haus, Englisch: house
	// entry4: Deutsch: Fahrrad, Englisch: bicycle
	// entry5: Deutsch: Auto, Englisch: car
}

/*
ExampleEntry_usage demonstriert, wie man Objekte vom Typ Entry verwendet.
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
	// entry1: Deutsch: Pflanze, Englisch: flower
	// entry1.De: Pflanze
	// entry1.En: flower
	//
	// entry1 nach der Änderung von entry1.En:
	// entry1: Deutsch: Pflanze, Englisch: plant
	// entry1.De: Pflanze
	// entry1.En: plant
}
