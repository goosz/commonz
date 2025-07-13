package commonz_test

import (
	"fmt"
	"reflect"

	"github.com/goosz/commonz"
)

// ExampleTypeName demonstrates how to use TypeName for debugging and logging.
func ExampleTypeName() {
	// Define a sample struct
	type User struct {
		ID   int
		Name string
	}

	// Create some example types
	users := []*User{}
	userMap := map[string]*User{}
	userChan := make(chan *User)

	// Get type names for debugging/logging
	fmt.Printf("Users slice: %s\n", commonz.TypeName(reflect.TypeOf(users)))
	fmt.Printf("User map: %s\n", commonz.TypeName(reflect.TypeOf(userMap)))
	fmt.Printf("User channel: %s\n", commonz.TypeName(reflect.TypeOf(userChan)))

	// Compare with reflect.Type.String()
	fmt.Printf("reflect.String() for users: %s\n", reflect.TypeOf(users).String())

	// Output:
	// Users slice: []*github.com/goosz/commonz_test.User
	// User map: map[string]*github.com/goosz/commonz_test.User
	// User channel: chan *github.com/goosz/commonz_test.User
	// reflect.String() for users: []*commonz_test.User
}

// ExampleTypeName_basicTypes shows TypeName with basic Go types.
func ExampleTypeName_basicTypes() {
	fmt.Printf("int: %s\n", commonz.TypeName(reflect.TypeOf(42)))
	fmt.Printf("string: %s\n", commonz.TypeName(reflect.TypeOf("hello")))
	fmt.Printf("bool: %s\n", commonz.TypeName(reflect.TypeOf(true)))
	fmt.Printf("float64: %s\n", commonz.TypeName(reflect.TypeOf(3.14)))

	// Output:
	// int: int
	// string: string
	// bool: bool
	// float64: float64
}

// ExampleTypeName_compositeTypes demonstrates TypeName with composite types.
func ExampleTypeName_compositeTypes() {
	// Arrays and slices
	intArray := [3]int{1, 2, 3}
	stringSlice := []string{"a", "b", "c"}

	// Maps
	stringIntMap := map[string]int{"a": 1, "b": 2}

	// Pointers
	intPtr := new(int)

	// Functions
	funcVar := func(x int, y string) error { return nil }

	fmt.Printf("Array: %s\n", commonz.TypeName(reflect.TypeOf(intArray)))
	fmt.Printf("Slice: %s\n", commonz.TypeName(reflect.TypeOf(stringSlice)))
	fmt.Printf("Map: %s\n", commonz.TypeName(reflect.TypeOf(stringIntMap)))
	fmt.Printf("Pointer: %s\n", commonz.TypeName(reflect.TypeOf(intPtr)))
	fmt.Printf("Function: %s\n", commonz.TypeName(reflect.TypeOf(funcVar)))

	// Output:
	// Array: [3]int
	// Slice: []string
	// Map: map[string]int
	// Pointer: *int
	// Function: func(int, string) error
}

// ExampleTypeName_channels shows TypeName with different channel types.
func ExampleTypeName_channels() {
	bidirChan := make(chan int)
	recvChan := make(<-chan string)
	sendChan := make(chan<- bool)

	fmt.Printf("Bidirectional channel: %s\n", commonz.TypeName(reflect.TypeOf(bidirChan)))
	fmt.Printf("Receive-only channel: %s\n", commonz.TypeName(reflect.TypeOf(recvChan)))
	fmt.Printf("Send-only channel: %s\n", commonz.TypeName(reflect.TypeOf(sendChan)))

	// Output:
	// Bidirectional channel: chan int
	// Receive-only channel: <-chan string
	// Send-only channel: chan<- bool
}
