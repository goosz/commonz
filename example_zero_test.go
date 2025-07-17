package commonz_test

import (
	"fmt"

	"github.com/goosz/commonz"
)

func ExampleZero_int() {
	fmt.Println(commonz.Zero[int]())
	// Output: 0
}

func ExampleZero_string() {
	fmt.Println(commonz.Zero[string]())
	// Output:
}

func ExampleZero_struct() {
	type S struct{ X int }
	fmt.Println(commonz.Zero[S]())
	// Output: {0}
}

func ExampleZero_slice() {
	fmt.Println(commonz.Zero[[]int]())
	// Output: []
}

func ExampleZero_map() {
	fmt.Println(commonz.Zero[map[string]int]())
	// Output: map[]
}
