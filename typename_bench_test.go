package commonz_test

import (
	"reflect"
	"testing"

	"github.com/goosz/commonz"
)

// BenchmarkTypeName_basicTypes benchmarks TypeName with basic Go types
func BenchmarkTypeName_basicTypes(b *testing.B) {
	types := []reflect.Type{
		reflect.TypeOf(42),               // int
		reflect.TypeOf("hello"),          // string
		reflect.TypeOf(true),             // bool
		reflect.TypeOf(3.14),             // float64
		reflect.TypeOf([]int{}),          // slice
		reflect.TypeOf(map[string]int{}), // map
		reflect.TypeOf(new(int)),         // pointer
		reflect.TypeOf(make(chan int)),   // channel
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range types {
			commonz.TypeName(t)
		}
	}
}

// BenchmarkTypeName_complexTypes benchmarks TypeName with complex nested types
func BenchmarkTypeName_complexTypes(b *testing.B) {
	types := []reflect.Type{
		reflect.TypeOf([][]map[string]*[]int{}),
		reflect.TypeOf(map[string]*[]map[int]chan<- bool{}),
		reflect.TypeOf(func([][]string, ...map[int]*[]bool) (chan<- *[]int, error) { return nil, nil }),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range types {
			commonz.TypeName(t)
		}
	}
}

// BenchmarkTypeName_structs benchmarks TypeName with struct types
func BenchmarkTypeName_structs(b *testing.B) {
	type SimpleStruct struct {
		Field int
	}

	type ComplexStruct struct {
		Simple  SimpleStruct
		Pointer *SimpleStruct
		Slice   []SimpleStruct
		Map     map[string]SimpleStruct
	}

	types := []reflect.Type{
		reflect.TypeOf(SimpleStruct{}),
		reflect.TypeOf(ComplexStruct{}),
		reflect.TypeOf([]ComplexStruct{}),
		reflect.TypeOf(map[string]*ComplexStruct{}),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range types {
			commonz.TypeName(t)
		}
	}
}

// BenchmarkTypeName_generics benchmarks TypeName with generic types
func BenchmarkTypeName_generics(b *testing.B) {
	type GenericStruct[T any] struct {
		Field T
	}

	type User struct {
		ID   int
		Name string
	}

	types := []reflect.Type{
		reflect.TypeOf(GenericStruct[int]{}),
		reflect.TypeOf(GenericStruct[User]{}),
		reflect.TypeOf([]GenericStruct[User]{}),
		reflect.TypeOf(map[string]GenericStruct[User]{}),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, t := range types {
			commonz.TypeName(t)
		}
	}
}

// BenchmarkTypeName_depthLimit benchmarks TypeName with types that hit the depth limit
func BenchmarkTypeName_depthLimit(b *testing.B) {
	// Create a deeply nested type that will hit the maxTypeDepth limit
	typ := reflect.TypeOf(map[int]map[string]map[bool]map[float64]map[string]map[int]map[bool]map[string]int{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		commonz.TypeName(typ)
	}
}

// BenchmarkTypeName_vsReflect benchmarks TypeName vs reflect.Type.String()
func BenchmarkTypeName_vsReflect(b *testing.B) {
	type User struct {
		ID   int
		Name string
	}

	typ := reflect.TypeOf([]*User{})

	b.Run("TypeName", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = commonz.TypeName(typ)
		}
	})

	b.Run("ReflectString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = typ.String()
		}
	})
}
