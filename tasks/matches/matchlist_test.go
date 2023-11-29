package matches

import (
	"fmt"
	"strings"
)

func ExampleMatchList_String() {
	ml := NewMatchList(
		NewMatch("Fortuna Fürth", "FC Frankfurt", NewScore(1, 2)),
		NewMatch("Dynamo Dortmund", "Concordia Cottbus", NewScore(3, 3)),
		NewMatch("Fortuna Fürth", "Dynamo Dortmund", NewScore(2, 1)),
		NewMatch("FC Frankfurt", "Concordia Cottbus", NewScore(1, 2)),
		NewMatch("Fortuna Fürth", "Concordia Cottbus", NewScore(2, 2)),
		NewMatch("FC Frankfurt", "Dynamo Dortmund", NewScore(3, 2)),
	)

	fmt.Println(ml)

	// Output:
	// Fortuna Fürth - *FC Frankfurt*: 1:2
	// Dynamo Dortmund - Concordia Cottbus: 3:3
	// *Fortuna Fürth* - Dynamo Dortmund: 2:1
	// FC Frankfurt - *Concordia Cottbus*: 1:2
	// Fortuna Fürth - Concordia Cottbus: 2:2
	// *FC Frankfurt* - Dynamo Dortmund: 3:2
}

func ExampleMatchList_Add() {
	ml := NewMatchList()

	// Spieltag 1
	ml.Add(NewMatch("Fortuna Fürth", "FC Frankfurt", NewScore(1, 2)))
	ml.Add(NewMatch("Dynamo Dortmund", "Concordia Cottbus", NewScore(3, 3)))

	// Spieltag 2
	ml.Add(NewMatch("Fortuna Fürth", "Dynamo Dortmund", NewScore(2, 1)))
	ml.Add(NewMatch("FC Frankfurt", "Concordia Cottbus", NewScore(1, 2)))

	// Spieltag 3
	ml.Add(NewMatch("Fortuna Fürth", "Concordia Cottbus", NewScore(2, 2)))
	ml.Add(NewMatch("FC Frankfurt", "Dynamo Dortmund", NewScore(3, 2)))

	fmt.Println(ml)

	// Output:
	// Fortuna Fürth - *FC Frankfurt*: 1:2
	// Dynamo Dortmund - Concordia Cottbus: 3:3
	// *Fortuna Fürth* - Dynamo Dortmund: 2:1
	// FC Frankfurt - *Concordia Cottbus*: 1:2
	// Fortuna Fürth - Concordia Cottbus: 2:2
	// *FC Frankfurt* - Dynamo Dortmund: 3:2
}

func ExampleMatchList_Points() {
	ml := NewMatchList(
		NewMatch("Fortuna Fürth", "FC Frankfurt", NewScore(1, 2)),
		NewMatch("Dynamo Dortmund", "Concordia Cottbus", NewScore(3, 3)),
		NewMatch("Fortuna Fürth", "Dynamo Dortmund", NewScore(2, 1)),
		NewMatch("FC Frankfurt", "Concordia Cottbus", NewScore(1, 2)),
		NewMatch("Fortuna Fürth", "Concordia Cottbus", NewScore(2, 2)),
		NewMatch("FC Frankfurt", "Dynamo Dortmund", NewScore(3, 2)),
	)

	fmt.Printf("Fortuna Fürth: %v\n", ml.Points("Fortuna Fürth"))
	fmt.Printf("FC Frankfurt: %v\n", ml.Points("FC Frankfurt"))
	fmt.Printf("Dynamo Dortmund: %v\n", ml.Points("Dynamo Dortmund"))
	fmt.Printf("Concordia Cottbus: %v\n", ml.Points("Concordia Cottbus"))

	// Output:
	// Fortuna Fürth: 4
	// FC Frankfurt: 6
	// Dynamo Dortmund: 1
	// Concordia Cottbus: 5
}

func ExampleMatchList_GoalDiff() {
	ml := NewMatchList(
		NewMatch("Fortuna Fürth", "FC Frankfurt", NewScore(1, 2)),
		NewMatch("Dynamo Dortmund", "Concordia Cottbus", NewScore(3, 3)),
		NewMatch("Fortuna Fürth", "Dynamo Dortmund", NewScore(2, 1)),
		NewMatch("FC Frankfurt", "Concordia Cottbus", NewScore(1, 2)),
		NewMatch("Fortuna Fürth", "Concordia Cottbus", NewScore(2, 2)),
		NewMatch("FC Frankfurt", "Dynamo Dortmund", NewScore(3, 2)),
	)

	fmt.Printf("Fortuna Fürth: %v\n", ml.GoalDiff("Fortuna Fürth"))
	fmt.Printf("FC Frankfurt: %v\n", ml.GoalDiff("FC Frankfurt"))
	fmt.Printf("Dynamo Dortmund: %v\n", ml.GoalDiff("Dynamo Dortmund"))
	fmt.Printf("Concordia Cottbus: %v\n", ml.GoalDiff("Concordia Cottbus"))

	// Output:
	// Fortuna Fürth: 0
	// FC Frankfurt: 1
	// Dynamo Dortmund: -2
	// Concordia Cottbus: 1
}

func ExampleMatchList_Teams() {
	ml := NewMatchList(
		NewMatch("Fortuna Fürth", "FC Frankfurt", NewScore(1, 2)),
		NewMatch("Dynamo Dortmund", "Concordia Cottbus", NewScore(3, 3)),
		NewMatch("Fortuna Fürth", "Dynamo Dortmund", NewScore(2, 1)),
		NewMatch("FC Frankfurt", "Concordia Cottbus", NewScore(1, 2)),
		NewMatch("Fortuna Fürth", "Concordia Cottbus", NewScore(2, 2)),
		NewMatch("FC Frankfurt", "Dynamo Dortmund", NewScore(3, 2)),
	)

	fmt.Println(ml.Teams())

	// Output:
	// [Concordia Cottbus Dynamo Dortmund FC Frankfurt Fortuna Fürth]
}

func ExampleMatchList_Compare() {
	ml := NewMatchList(
		NewMatch("Fortuna Fürth", "FC Frankfurt", NewScore(1, 2)),
		NewMatch("Dynamo Dortmund", "Concordia Cottbus", NewScore(3, 3)),
		NewMatch("Fortuna Fürth", "Dynamo Dortmund", NewScore(2, 1)),
		NewMatch("FC Frankfurt", "Concordia Cottbus", NewScore(1, 2)),
		NewMatch("Fortuna Fürth", "Concordia Cottbus", NewScore(2, 2)),
		NewMatch("FC Frankfurt", "Dynamo Dortmund", NewScore(3, 2)),
	)

	fmt.Println(ml.Compare("Fortuna Fürth", "FC Frankfurt"))
	fmt.Println(ml.Compare("FC Frankfurt", "Fortuna Fürth"))
	fmt.Println(ml.Compare("FC Frankfurt", "FC Frankfurt"))

	// Output:
	// 1
	// -1
	// 0
}

func ExampleMatchList_Table() {
	ml := NewMatchList(
		NewMatch("Fortuna Fürth", "FC Frankfurt", NewScore(1, 2)),
		NewMatch("Dynamo Dortmund", "Concordia Cottbus", NewScore(3, 3)),
		NewMatch("Fortuna Fürth", "Dynamo Dortmund", NewScore(2, 1)),
		NewMatch("FC Frankfurt", "Concordia Cottbus", NewScore(1, 2)),
		NewMatch("Fortuna Fürth", "Concordia Cottbus", NewScore(2, 2)),
		NewMatch("FC Frankfurt", "Dynamo Dortmund", NewScore(3, 2)),
	)

	table := ml.Table()
	fmt.Println(strings.Join(table, "\n"))

	// Output:
	// 1 FC Frankfurt 6 1
	// 2 Concordia Cottbus 5 1
	// 3 Fortuna Fürth 4 0
	// 4 Dynamo Dortmund 1 -2
}
