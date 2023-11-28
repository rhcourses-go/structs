package dict1

import "fmt"

// ExampleDict_definition demonstriert, wie neue Objekte vom Typ Dict erzeugt werden können.
// Anmerkung: Die Ausdrücke in diesem Beispiel sind noch recht kompliziert in der Verwendung.
// Im Package dict2 werden wir sehen, wie man diese Komplexität für den Client-Programmierer
// verringern kann.
func ExampleDict_definition() {
	// Definition eines leeren Wörterbuch-Objekts:
	dict1 := Dict{}

	// Definition eines Wörterbuch-Objekts mit einigen Einträgen:
	dict2 := Dict{Entries: []Entry{
		{De: "Sprache", En: "language"},
		{De: "kompliziert", En: "complicated"},
		{De: "einfach", En: "simple"},
	}}

	// Ausgabe der Wörterbücher:
	fmt.Printf("dict1: %+v\n", dict1)
	fmt.Printf("dict2: %+v\n", dict2)

	// Output:
	// dict1: {Entries:[]}
	// dict2: {Entries:[{De:Sprache En:language} {De:kompliziert En:complicated} {De:einfach En:simple}]}
}

// ExampleDict_usage demonstriert, wie ein Objekt vom Typ `Dict` ausgelesen und manipuliert werden kann.
func ExampleDict_usage() {
	// Definition eines leeren Wörterbuch-Objekts:
	dict1 := Dict{}

	// Erzeugen einiger neuen Einträge:
	entry1 := Entry{De: "Baum", En: "tree"}
	entry2 := Entry{De: "Wald", En: "forest"}
	entry3 := Entry{De: "Wald", En: "wood"}
	entry4 := Entry{De: "Holz", En: "wood"}

	// Hinzufügen der Einträge zu `dict1`:
	dict1.Entries = append(dict1.Entries, entry1, entry2, entry3, entry4)

	// Ausgabe des Wörterbuchs:
	fmt.Printf("dict1: %+v\n", dict1)

	// Ausgabe aller Übersetzungen des Wortes "Wald":
	for _, entry := range dict1.Entries {
		if entry.De == "Wald" {
			fmt.Println(entry.En)
		}
	}

	// Bestimmen und Ausgeben aller deutschen Wörter, die mit "wood" übersetzt werden können:
	wood_words := []string{}
	for _, entry := range dict1.Entries {
		if entry.En == "wood" {
			wood_words = append(wood_words, entry.De)
		}
	}
	fmt.Println(wood_words)

	// Output:
	// dict1: {Entries:[{De:Baum En:tree} {De:Wald En:forest} {De:Wald En:wood} {De:Holz En:wood}]}
	// forest
	// wood
	// [Wald Holz]

}
