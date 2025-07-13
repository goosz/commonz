package commonz_test

import (
	"reflect"
	"testing"

	. "github.com/goosz/commonz"
	"github.com/stretchr/testify/assert"
)

type Struct struct {
	Field int
}

type StructWithGenericArg[T any] struct {
	Field T
}

type Interface interface {
	Function()
}

type InterfaceWithGenericArg[T any] interface {
	Function() T
}

func TestTypeName(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "bool",
			input:    true,
			expected: "bool",
		},
		{
			name:     "int",
			input:    42,
			expected: "int",
		},
		{
			name:     "string",
			input:    "hello",
			expected: "string",
		},
		{
			name:     "array",
			input:    [3]int{1, 2, 3},
			expected: "[3]int",
		},
		{
			name:     "slice",
			input:    []string{"a", "b"},
			expected: "[]string",
		},
		{
			name:     "slice of empty interface (a.k.a. any)",
			input:    []any{},
			expected: "[]interface {}",
		},
		{
			name:     "slice of struct",
			input:    []Struct{},
			expected: "[]github.com/goosz/commonz_test.Struct",
		},
		{
			name:     "slice of struct with generic arg",
			input:    []StructWithGenericArg[Struct]{},
			expected: "[]github.com/goosz/commonz_test.StructWithGenericArg[github.com/goosz/commonz_test.Struct]",
		},
		{
			name:     "slice of interface",
			input:    []Interface{},
			expected: "[]github.com/goosz/commonz_test.Interface",
		},
		{
			name:     "slice of interface with generic arg",
			input:    []InterfaceWithGenericArg[Interface]{},
			expected: "[]github.com/goosz/commonz_test.InterfaceWithGenericArg[github.com/goosz/commonz_test.Interface]",
		},
		{
			name:     "map",
			input:    map[string]int{"a": 1},
			expected: "map[string]int",
		},
		{
			name:     "map of struct",
			input:    map[Struct]Struct{},
			expected: "map[github.com/goosz/commonz_test.Struct]github.com/goosz/commonz_test.Struct",
		},
		{
			name:     "map of struct with generic arg",
			input:    map[StructWithGenericArg[Struct]]StructWithGenericArg[Struct]{},
			expected: "map[github.com/goosz/commonz_test.StructWithGenericArg[github.com/goosz/commonz_test.Struct]]github.com/goosz/commonz_test.StructWithGenericArg[github.com/goosz/commonz_test.Struct]",
		},
		{
			name:     "map of interface",
			input:    map[Interface]Interface{},
			expected: "map[github.com/goosz/commonz_test.Interface]github.com/goosz/commonz_test.Interface",
		},
		{
			name:     "map of interface with generic arg",
			input:    map[InterfaceWithGenericArg[Interface]]InterfaceWithGenericArg[Interface]{},
			expected: "map[github.com/goosz/commonz_test.InterfaceWithGenericArg[github.com/goosz/commonz_test.Interface]]github.com/goosz/commonz_test.InterfaceWithGenericArg[github.com/goosz/commonz_test.Interface]",
		},
		{
			name:     "channel",
			input:    make(chan int),
			expected: "chan int",
		},
		{
			name:     "receive channel",
			input:    make(<-chan int),
			expected: "<-chan int",
		},
		{
			name:     "send channel",
			input:    make(chan<- int),
			expected: "chan<- int",
		},
		{
			name:     "channel of struct",
			input:    make(chan Struct),
			expected: "chan github.com/goosz/commonz_test.Struct",
		},
		{
			name:     "pointer",
			input:    new(int),
			expected: "*int",
		},
		{
			name:     "pointer to struct",
			input:    new(Struct),
			expected: "*github.com/goosz/commonz_test.Struct",
		},
		{
			name:     "function",
			input:    func(int, int) {},
			expected: "func(int, int)",
		},
		{
			name:     "function with variadic argument",
			input:    func(int, ...int) {},
			expected: "func(int, ...int)",
		},
		{
			name:     "function with return type",
			input:    func(int, int) (int, error) { return 0, nil },
			expected: "func(int, int) (int, error)",
		},
		{
			name:     "function with struct arguments and return type",
			input:    func(Struct, ...Struct) Struct { return Struct{} },
			expected: "func(github.com/goosz/commonz_test.Struct, ...github.com/goosz/commonz_test.Struct) github.com/goosz/commonz_test.Struct",
		},
		{
			name:     "struct",
			input:    Struct{},
			expected: "github.com/goosz/commonz_test.Struct",
		},
		{
			name:     "struct with generic arg",
			input:    StructWithGenericArg[Struct]{},
			expected: "github.com/goosz/commonz_test.StructWithGenericArg[github.com/goosz/commonz_test.Struct]",
		},
		{
			name:     "[]error",
			input:    []error{},
			expected: "[]error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, TypeName(reflect.TypeOf(tt.input)))
		})
	}
}

func TestTypeNameNil(t *testing.T) {
	assert.Equal(t, "<nil>", TypeName(nil))
}

func TestTypeNameNestedOverflow(t *testing.T) {
	typ := reflect.TypeOf(map[int]map[string]map[int]map[string]int{})
	assert.Equal(t, "map[int]map[string]map[int]map[<...>]<...>", TypeNameWithDepth(typ, 4))
}
