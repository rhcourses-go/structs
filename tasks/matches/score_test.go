package matches

import "fmt"

func ExampleScore() {
	s := NewScore(1, 2)

	fmt.Println(s)

	// Output:
	// 1:2
}

func ExampleScore_Result() {
	s1 := NewScore(1, 2)
	s2 := NewScore(2, 1)
	s3 := NewScore(2, 2)

	fmt.Println(s1.Result())
	fmt.Println(s2.Result())
	fmt.Println(s3.Result())

	// Output:
	// 2
	// 1
	// 0
}
