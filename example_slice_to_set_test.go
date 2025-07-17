package commonz_test

import (
	"fmt"

	"github.com/goosz/commonz"
)

func ExampleSliceToSet_basic() {
	slice := []string{"apple", "banana", "apple", "cherry"}
	set, err := commonz.SliceToSet(slice, false)
	fmt.Println("err:", err)
	for k := range set {
		fmt.Println(k)
	}
	// Unordered output:
	// err: <nil>
	// apple
	// banana
	// cherry
}

func ExampleSliceToSet_strict() {
	slice := []int{1, 2, 2, 3}
	_, err := commonz.SliceToSet(slice, true)
	fmt.Println("err:", err)
	// Output:
	// err: duplicate element: 2
}
