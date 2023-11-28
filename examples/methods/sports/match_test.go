package sports

import "fmt"

func ExampleMatch_SetHomeTeam() {
	// Erzeuge ein Match-Objekt:
	m := Match{Home: "FC Freiburg", Visitors: "Borussia Bremen", Location: "Mannheim", Result: 0}

	// m ausgeben:
	fmt.Println(m)

	// Schreibe einen neuen Namen ins Home-Feld mittels der Methode SetHomeTeam:
	m.SetHomeTeam("Bayer 04 Bayern")

	// m ausgeben:
	fmt.Println(m)

	// Output:
	// {FC Freiburg Borussia Bremen Mannheim 0}
	// {Bayer 04 Bayern Borussia Bremen Mannheim 0}
}

func ExampleMatch_SetVisitorTeam() {
	// Erzeuge ein Match-Objekt:
	m := Match{Home: "FC Freiburg", Visitors: "Borussia Bremen", Location: "Mannheim", Result: 0}

	// m ausgeben:
	fmt.Println(m)

	// Schreibe einen Namen ins Visitors-Feld mittels der Methode SetVisitorTeam:
	m.SetVisitorTeam("SC Schalke")

	// m ausgeben:
	fmt.Println(m)

	// Output:
	// {FC Freiburg Borussia Bremen Mannheim 0}
	// {FC Freiburg SC Schalke Mannheim 0}
}

func ExampleMatch_SetLocation() {
	// Erzeuge ein Match-Objekt mit leerem Heim- und leerem Location-Feld:
	m := Match{Home: "", Visitors: "Werder Würzburg", Location: "", Result: 0}
	fmt.Println(m)

	// Schreibe einen Namen ins Location-Feld von m1 mittels der Methode SetLocation.
	// Da das Home-Feld leer ist, erwarten wir, dass auch dies gesetzt wird.
	m.SetLocation("Union Ulm")
	fmt.Println(m)

	// Schreibe die Location erneut.
	// Da das Home-Feld nun nicht mehr leer ist,
	// erwarten wir, dass dies nicht mehr verändert wird.
	m.SetLocation("Ulm")
	fmt.Println(m)

	// Output:
	//{ Werder Würzburg  0}
	//{Union Ulm Werder Würzburg Union Ulm 0}
	//{Union Ulm Werder Würzburg Ulm 0}
}
