package commonz

import "fmt"

// SliceToSet converts a slice of comparable elements to a set (map[T]struct{}).
// If strict is true, it returns an error on duplicates; otherwise, it silently deduplicates.
//
// The type parameter T must be comparable, as required for map keys in Go. Note that:
//   - If T is a struct containing non-comparable fields (such as slices, maps, or functions),
//     this function will not compile.
//   - If T is an interface type, the function will compile, because all interfaces are comparable.
//     However, if any element in the slice has a dynamic type that is not comparable (e.g., a struct
//     containing a slice), inserting or comparing such values as map keys will cause a runtime panic.
//
// Example (will not compile):
//
//	type NotComparable struct { Data []int }
//	_, _ = SliceToSet([]NotComparable{{Data: []int{1,2}}}, false)
//
// Example (compiles, but panics at runtime):
//
//	type Demo interface{ Foo() }
//	type NotComparableStruct struct{ Data []int }
//	func (NotComparableStruct) Foo() {}
//	slice := []Demo{NotComparableStruct{Data: []int{1,2}}}
//	_, _ = SliceToSet(slice, false) // panics at runtime when using map key
func SliceToSet[T comparable](slice []T, strict bool) (map[T]struct{}, error) {
	set := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		if _, exists := set[v]; exists && strict {
			return nil, fmt.Errorf("duplicate element: %v", v)
		}
		set[v] = struct{}{}
	}
	return set, nil
}
