package dict2

import "fmt"

// ExampleDict_definition demonstriert, wie neue Objekte vom Typ Dict erzeugt werden können.
// Im Gegensatz zum Beispiel im Package dict1 wird hier der Konstruktor `NewDict` verwendet.
// Außerdem ist auch hier die Ausgabe verändert, weil es String-Methoden gibt.
func ExampleDict_definition() {
	// Definition eines leeren Wörterbuch-Objekts:
	dict1 := Dict{}

	// Definition eines Wörterbuch-Objekts mit einigen Einträgen:
	dict2 := NewDict(
		Entry{De: "Sprache", En: "language"},
		Entry{De: "kompliziert", En: "complicated"},
		Entry{De: "einfach", En: "simple"},
	)

	// Ausgabe der Wörterbücher:
	fmt.Printf("dict1:\n%v\n", dict1)
	fmt.Printf("dict2:\n%v\n", dict2)

	// Output:
	// dict1:
	// keine Einträge
	// dict2:
	// Deutsch: Sprache, Englisch: language
	// Deutsch: kompliziert, Englisch: complicated
	// Deutsch: einfach, Englisch: simple
}

// ExampleDict_usage demonstriert, wie ein Objekt vom Typ `Dict` ausgelesen und manipuliert werden kann.
// Im Gegensatz zum Beispiel im Package dict1 wird hier die Methode `Append` verwendet.
// Außerdem ist die String-Ausgabe per Println einfacher, weil es eine Methode `String` gibt
// und die Abfragen geschehen auch über Methoden.
func ExampleDict_usage() {
	// Definition eines leeren Wörterbuch-Objekts:
	dict1 := Dict{}

	// Erzeugen einiger neuen Einträge:
	entry1 := Entry{De: "Baum", En: "tree"}
	entry2 := Entry{De: "Wald", En: "forest"}
	entry3 := Entry{De: "Wald", En: "wood"}
	entry4 := Entry{De: "Holz", En: "wood"}

	// Hinzufügen der Einträge zu `dict1`:
	dict1.Append(entry1, entry2, entry3, entry4)

	// Ausgabe des Wörterbuchs:
	fmt.Printf("dict1:\n%+v\n", dict1)

	// Ausgabe aller englischen Übersetzungen des deutschen Wortes "Wald":
	fmt.Println(dict1.LookupDe("Wald"))

	// Ausgabe aller deutschen Übersetzungen des englischen Wortes "wood:
	fmt.Println(dict1.LookupEn("wood"))

	// Output:
	// dict1:
	// Deutsch: Baum, Englisch: tree
	// Deutsch: Wald, Englisch: forest
	// Deutsch: Wald, Englisch: wood
	// Deutsch: Holz, Englisch: wood
	// [forest wood]
	// [Wald Holz]

}
