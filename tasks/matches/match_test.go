package matches

import "fmt"

func ExampleMatch_HomeName() {
	m1 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(1, 2))
	m2 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(2, 1))
	m3 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(2, 2))

	fmt.Println(m1.HomeName())
	fmt.Println(m2.HomeName())
	fmt.Println(m3.HomeName())

	// Output:
	// FC Freiburg
	// *FC Freiburg*
	// FC Freiburg
}

func ExampleMatch_VisitorName() {
	m1 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(1, 2))
	m2 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(2, 1))
	m3 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(2, 2))

	fmt.Println(m1.VisitorName())
	fmt.Println(m2.VisitorName())
	fmt.Println(m3.VisitorName())

	// Output:
	// *Borussia Bremen*
	// Borussia Bremen
	// Borussia Bremen
}

func ExampleMatch_String() {
	m := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(1, 2))

	fmt.Println(m)

	// Output:
	// FC Freiburg - *Borussia Bremen*: 1:2
}

func ExampleMatch_Winner() {
	m1 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(1, 2))
	m2 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(2, 1))
	m3 := NewMatch("FC Freiburg", "Borussia Bremen", NewScore(2, 2))

	fmt.Println(m1.Winner())
	fmt.Println(m2.Winner())
	fmt.Println(m3.Winner())

	// Output:
	// Borussia Bremen
	// FC Freiburg
	// unentschieden
}
