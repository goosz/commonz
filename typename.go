package commonz

import (
	"fmt"
	"reflect"
)

// Prevent stack overflow with deeply nested types.
// This should be more than enough for reasonable uses.
//
// Most real-world Go code rarely exceeds 4-5 levels of nesting, making 8 a safe
// upper bound that prevents stack overflow while preserving meaningful type information.
const maxTypeDepth = 8

// TypeName returns a human-readable string representation of a Go type.
// Unlike [reflect.Type]'s String() implementation, this function includes
// package paths for structs and interfaces when available, making type
// identification clearer in multi-package contexts.
//
// The function handles all Go types including basic types, composite types,
// function types, generic types, and channels. See the examples for usage.
func TypeName(t reflect.Type) string {
	// Handle nil type
	if t == nil {
		return "<nil>"
	}
	return typeNameWithDepth(t, maxTypeDepth)
}

func typeNameForFunc(t reflect.Type, maxDepth int) string {
	paramTypes := make([]string, t.NumIn())
	for i := range t.NumIn() {
		if t.IsVariadic() && i == t.NumIn()-1 {
			paramTypes[i] = "..." + typeNameWithDepth(t.In(i).Elem(), maxDepth-1)
		} else {
			paramTypes[i] = typeNameWithDepth(t.In(i), maxDepth-1)
		}
	}
	params := ""
	if len(paramTypes) > 0 {
		params = paramTypes[0]
		for i := 1; i < len(paramTypes); i++ {
			params += ", " + paramTypes[i]
		}
	}
	if t.NumOut() > 0 {
		returnTypes := make([]string, t.NumOut())
		for i := range t.NumOut() {
			returnTypes[i] = typeNameWithDepth(t.Out(i), maxDepth-1)
		}
		var returns string
		if len(returnTypes) > 1 {
			returns += "("
		}
		returns += returnTypes[0]
		for i := 1; i < len(returnTypes); i++ {
			returns += ", " + returnTypes[i]
		}
		if len(returnTypes) > 1 {
			returns += ")"
		}
		return fmt.Sprintf("func(%s) %s", params, returns)
	}
	return fmt.Sprintf("func(%s)", params)
}

func typeNameWithDepth(t reflect.Type, maxDepth int) string {
	// Prevent stack overflows with deeply nested types
	if maxDepth == 0 {
		return "<...>"
	}

	switch t.Kind() {
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", t.Len(), typeNameWithDepth(t.Elem(), maxDepth-1))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", typeNameWithDepth(t.Elem(), maxDepth-1))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s",
			typeNameWithDepth(t.Key(), maxDepth-1),
			typeNameWithDepth(t.Elem(), maxDepth-1))
	case reflect.Chan:
		dir := ""
		switch t.ChanDir() {
		case reflect.RecvDir:
			dir = "<-chan "
		case reflect.SendDir:
			dir = "chan<- "
		default:
			dir = "chan "
		}
		return dir + typeNameWithDepth(t.Elem(), maxDepth-1)
	case reflect.Ptr:
		return "*" + typeNameWithDepth(t.Elem(), maxDepth-1)
	case reflect.Struct, reflect.Interface:
		if t.PkgPath() == "" {
			return t.String()
		}
		return fmt.Sprintf("%s.%s", t.PkgPath(), t.Name())
	case reflect.Func:
		return typeNameForFunc(t, maxDepth)
	default:
		return t.String() // Use reflect.Type's own String() implementation for everything else
	}
}
