package functions

import (
	"fmt"
)

// ExampleToColumns
func ExampleToColumns() {
	w := 40
	s := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.`

	got := ToColumns(w, s)

	fmt.Println(got)
	// Output:
	// Lorem ipsum dolor sit amet, consectetur
	// adipiscing elit, sed do eiusmod tempor
	// incididunt ut labore et dolore magna
	// aliqua. Ut enim ad minim veniam, quis
	// nostrud exercitation ullamco laboris
	// nisi ut aliquip ex ea commodo consequat.
}
