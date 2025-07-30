package commonz_test

import (
	"testing"

	"github.com/goosz/commonz"
	"github.com/stretchr/testify/require"
)

var initCaller = commonz.GetCaller(commonz.CurrentCaller)

func TestParseCallerInfo(t *testing.T) {
	tests := []struct {
		name     string
		fnName   string
		expected commonz.CallerInfo
	}{
		{
			name:   "empty string",
			fnName: "",
			expected: commonz.CallerInfo{
				Package:  "<unknown-package>",
				Function: "<unknown-function>",
			},
		},
		{
			name:   "string with no periods",
			fnName: "justafunctionname",
			expected: commonz.CallerInfo{
				Package:  "<unknown-package>",
				Function: "<unknown-function>",
			},
		},
		{
			name:   "exported function",
			fnName: "github.com/goosz/commonz_test.ExportedTestFunction",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "ExportedTestFunction",
			},
		},
		{
			name:   "unexported function",
			fnName: "github.com/goosz/commonz_test.unexportedTestFunction",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "unexportedTestFunction",
			},
		},
		{
			name:   "value receiver method",
			fnName: "github.com/goosz/commonz_test.ValueReceiverStruct.GetCallerInfo",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "ValueReceiverStruct.GetCallerInfo",
			},
		},
		{
			name:   "pointer receiver method",
			fnName: "github.com/goosz/commonz_test.(*PointerReceiverStruct).GetCallerInfo",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "(*PointerReceiverStruct).GetCallerInfo",
			},
		},
		{
			name:   "embedded value struct method",
			fnName: "github.com/goosz/commonz_test.EmbeddedValueStruct.GetCallerInfo",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "EmbeddedValueStruct.GetCallerInfo",
			},
		},
		{
			name:   "embedded pointer struct method",
			fnName: "github.com/goosz/commonz_test.(*EmbeddedPointerStruct).GetCallerInfo",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "(*EmbeddedPointerStruct).GetCallerInfo",
			},
		},
		{
			name:   "generic function",
			fnName: "github.com/goosz/commonz_test.GenericTestFunction[string]",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "GenericTestFunction[string]",
			},
		},
		{
			name:   "generic value receiver method",
			fnName: "github.com/goosz/commonz_test.GenericValueReceiver[string].GetCallerInfo",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "GenericValueReceiver[string].GetCallerInfo",
			},
		},
		{
			name:   "generic pointer receiver method",
			fnName: "github.com/goosz/commonz_test.(*GenericPointerReceiver[int]).GetCallerInfo",
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "(*GenericPointerReceiver[int]).GetCallerInfo",
			},
		},
		{
			name:   "foo.bar.com exported function",
			fnName: "foo.bar.com/mypackage.ExportedFunction",
			expected: commonz.CallerInfo{
				Package:  "foo.bar.com/mypackage",
				Function: "ExportedFunction",
			},
		},
		{
			name:   "foo.bar.com unexported function",
			fnName: "foo.bar.com/mypackage.unexportedFunction",
			expected: commonz.CallerInfo{
				Package:  "foo.bar.com/mypackage",
				Function: "unexportedFunction",
			},
		},
		{
			name:   "foo.bar.com value receiver method",
			fnName: "foo.bar.com/mypackage.ValueStruct.GetMethod",
			expected: commonz.CallerInfo{
				Package:  "foo.bar.com/mypackage",
				Function: "ValueStruct.GetMethod",
			},
		},
		{
			name:   "foo.bar.com pointer receiver method",
			fnName: "foo.bar.com/mypackage.(*PointerStruct).GetMethod",
			expected: commonz.CallerInfo{
				Package:  "foo.bar.com/mypackage",
				Function: "(*PointerStruct).GetMethod",
			},
		},
		{
			name:   "foo.bar.com generic function",
			fnName: "foo.bar.com/mypackage.GenericFunction[string]",
			expected: commonz.CallerInfo{
				Package:  "foo.bar.com/mypackage",
				Function: "GenericFunction[string]",
			},
		},
		{
			name:   "foo.bar.com generic value receiver method",
			fnName: "foo.bar.com/mypackage.GenericValueReceiver[string].GetMethod",
			expected: commonz.CallerInfo{
				Package:  "foo.bar.com/mypackage",
				Function: "GenericValueReceiver[string].GetMethod",
			},
		},
		{
			name:   "foo.bar.com generic pointer receiver method",
			fnName: "foo.bar.com/mypackage.(*GenericPointerReceiver[int]).GetMethod",
			expected: commonz.CallerInfo{
				Package:  "foo.bar.com/mypackage",
				Function: "(*GenericPointerReceiver[int]).GetMethod",
			},
		},
		{
			name:   "malformed pointer method - no closing parenthesis",
			fnName: "package.(*Type.GetMethod",
			expected: commonz.CallerInfo{
				Package:  "package",
				Function: "(*Type.GetMethod",
			},
		},
		{
			name:   "malformed pointer method - empty parentheses",
			fnName: "package.().GetMethod",
			expected: commonz.CallerInfo{
				Package:  "package",
				Function: "().GetMethod",
			},
		},
		{
			name:   "malformed pointer method - no opening parenthesis",
			fnName: "package.Type).GetMethod",
			expected: commonz.CallerInfo{
				Package:  "package",
				Function: "Type).GetMethod",
			},
		},
		{
			name:   "malformed pointer method - startPos out of bounds",
			fnName: "package.(*Type).GetMethod",
			expected: commonz.CallerInfo{
				Package:  "package",
				Function: "(*Type).GetMethod",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := commonz.ParseCallerInfo(tt.fnName)
			require.Equal(t, tt.expected, result, "ParseCallerInfo should return the correct components")
		})
	}
}

func TestCallerInfo_String(t *testing.T) {
	tests := []struct {
		name     string
		info     commonz.CallerInfo
		expected string
	}{
		{
			name: "simple function",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "TestFunction",
			},
			expected: "github.com/goosz/commonz_test.TestFunction",
		},
		{
			name: "pointer method",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "(*MyType).Method",
			},
			expected: "github.com/goosz/commonz_test.(*MyType).Method",
		},
		{
			name: "value method",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "MyType.Method",
			},
			expected: "github.com/goosz/commonz_test.MyType.Method",
		},
		{
			name: "generic function",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "GenericFunction[string]",
			},
			expected: "github.com/goosz/commonz_test.GenericFunction[string]",
		},
		{
			name: "generic value receiver method",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "Container[string].GetValue",
			},
			expected: "github.com/goosz/commonz_test.Container[string].GetValue",
		},
		{
			name: "generic pointer receiver method",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "(*Container[string]).SetValue",
			},
			expected: "github.com/goosz/commonz_test.(*Container[string]).SetValue",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.info.String()
			require.Equal(t, tt.expected, result, "CallerInfo.String() should return the correct representation")
		})
	}
}

func TestCallerInfo_IsUnknown(t *testing.T) {
	tests := []struct {
		name     string
		info     commonz.CallerInfo
		expected bool
	}{
		{
			name: "known caller",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "TestFunction",
			},
			expected: false,
		},
		{
			name: "unknown package",
			info: commonz.CallerInfo{
				Package:  "<unknown-package>",
				Function: "TestFunction",
			},
			expected: true,
		},
		{
			name: "unknown function",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "<unknown-function>",
			},
			expected: true,
		},
		{
			name: "both unknown",
			info: commonz.CallerInfo{
				Package:  "<unknown-package>",
				Function: "<unknown-function>",
			},
			expected: true,
		},
		{
			name: "empty strings",
			info: commonz.CallerInfo{
				Package:  "",
				Function: "",
			},
			expected: false,
		},
		{
			name: "partial unknown package",
			info: commonz.CallerInfo{
				Package:  "unknown-package",
				Function: "TestFunction",
			},
			expected: false,
		},
		{
			name: "partial unknown function",
			info: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "unknown-function",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.info.IsUnknown()
			require.Equal(t, tt.expected, result, "CallerInfo.IsUnknown() should return the correct value")
		})
	}
}

func TestGetCaller(t *testing.T) {
	tests := []struct {
		name     string
		depth    int
		expected commonz.CallerInfo
	}{
		{
			name:  "depth 0 - current function",
			depth: commonz.CurrentCaller,
			expected: commonz.CallerInfo{
				Package:  "github.com/goosz/commonz_test",
				Function: "TestGetCaller.func1",
			},
		},
		{
			name:  "depth 1 - testing package",
			depth: commonz.ParentCaller,
			expected: commonz.CallerInfo{
				Package:  "testing",
				Function: "tRunner",
			},
		},
		{
			name:  "depth 2 - runtime package",
			depth: commonz.GrandparentCaller,
			expected: commonz.CallerInfo{
				Package:  "runtime",
				Function: "goexit",
			},
		},
		{
			name:  "negative depth",
			depth: -1,
			expected: commonz.CallerInfo{
				Package:  "<unknown-package>",
				Function: "<unknown-function>",
			},
		},
		{
			name:  "large depth",
			depth: 100,
			expected: commonz.CallerInfo{
				Package:  "<unknown-package>",
				Function: "<unknown-function>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := commonz.GetCaller(tt.depth)
			require.Equal(t, tt.expected, result, "GetCaller should return the correct caller info")
		})
	}
}

// TestGetCallerFromExportedFunction tests GetCaller when called from an exported function
func TestGetCallerFromExportedFunction(t *testing.T) {
	// Test calling GetCaller from exported function
	caller := ExportedTestFunction()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "ExportedTestFunction",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from exported function should return correct caller info")
}

// TestGetCallerFromUnexportedFunction tests GetCaller when called from an unexported function
func TestGetCallerFromUnexportedFunction(t *testing.T) {
	// Test calling GetCaller from unexported function
	caller := unexportedTestFunction()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "unexportedTestFunction",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from unexported function should return correct caller info")
}

// TestGetCallerFromValueReceiver tests GetCaller when called from a value receiver method
func TestGetCallerFromValueReceiver(t *testing.T) {
	// Create instance and call value receiver method
	vs := ValueReceiverStruct{name: "test"}
	caller := vs.GetCallerInfo()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "ValueReceiverStruct.GetCallerInfo",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from value receiver method should return correct caller info")
}

// TestGetCallerFromPointerReceiver tests GetCaller when called from a pointer receiver method
func TestGetCallerFromPointerReceiver(t *testing.T) {
	// Create instance and call pointer receiver method
	ps := &PointerReceiverStruct{name: "test"}
	caller := ps.GetCallerInfo()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "(*PointerReceiverStruct).GetCallerInfo",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from pointer receiver method should return correct caller info")
}

// TestGetCallerFromEmbeddedValueStruct tests GetCaller when called from an embedded value struct method
func TestGetCallerFromEmbeddedValueStruct(t *testing.T) {
	// Create instance with embedded value receiver struct
	embedded := EmbeddedValueStruct{
		ValueReceiverStruct: ValueReceiverStruct{name: "test"},
		extraValue:          42,
	}

	// Test calling GetCaller from embedded value struct method
	caller := embedded.GetCallerInfo()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "ValueReceiverStruct.GetCallerInfo",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from embedded value struct method should return correct caller info")
}

// TestGetCallerFromEmbeddedPointerStruct tests GetCaller when called from an embedded pointer struct method
func TestGetCallerFromEmbeddedPointerStruct(t *testing.T) {
	// Create instance with embedded pointer receiver struct
	embedded := &EmbeddedPointerStruct{
		PointerReceiverStruct: &PointerReceiverStruct{name: "test"},
		extraValue:            42,
	}

	// Test calling GetCaller from embedded pointer struct method
	caller := embedded.GetCallerInfo()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "(*PointerReceiverStruct).GetCallerInfo",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from embedded pointer struct method should return correct caller info")
}

// TestGetCallerFromGenericFunction tests GetCaller when called from a generic simple function
func TestGetCallerFromGenericFunction(t *testing.T) {
	// Test calling GetCaller from generic simple function
	caller := GenericTestFunction("test")

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "GenericTestFunction[...]",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from generic simple function should return correct caller info")
}

// TestGetCallerFromGenericValueReceiver tests GetCaller when called from a value receiver method on a generic type
func TestGetCallerFromGenericValueReceiver(t *testing.T) {
	// Create instance of generic value receiver struct
	receiver := GenericValueReceiver[string]{value: "test"}

	// Test calling GetCaller from generic value receiver method
	caller := receiver.GetCallerInfo()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "GenericValueReceiver[...].GetCallerInfo",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from generic value receiver method should return correct caller info")
}

// TestGetCallerFromGenericPointerReceiver tests GetCaller when called from a pointer receiver method on a generic type
func TestGetCallerFromGenericPointerReceiver(t *testing.T) {
	// Create instance of generic pointer receiver struct
	receiver := &GenericPointerReceiver[int]{value: 42}

	// Test calling GetCaller from generic pointer receiver method
	caller := receiver.GetCallerInfo()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "(*GenericPointerReceiver[...]).GetCallerInfo",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from generic pointer receiver method should return correct caller info")
}

// TestGetCallerFromAnonymousFunction tests GetCaller when called from an anonymous function
func TestGetCallerFromAnonymousFunction(t *testing.T) {
	// Call GetCaller from an anonymous function
	caller := func() commonz.CallerInfo {
		return commonz.GetCaller(commonz.CurrentCaller)
	}()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "TestGetCallerFromAnonymousFunction.func1",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from anonymous function should return correct caller info")
}

// TestGetCallerFromAnonymousFunctionDepth1 tests GetCaller with depth 1 from an anonymous function
func TestGetCallerFromAnonymousFunctionDepth1(t *testing.T) {
	// Call GetCaller with depth 1 from an anonymous function
	caller := func() commonz.CallerInfo {
		return commonz.GetCaller(commonz.ParentCaller)
	}()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "TestGetCallerFromAnonymousFunctionDepth1",
	}

	require.Equal(t, expected, caller, "GetCaller(ParentCaller) from anonymous function should return correct caller info")
}

// TestGetCallerFromNestedAnonymousFunctions tests GetCaller with multiple levels of anonymous functions
func TestGetCallerFromNestedAnonymousFunctions(t *testing.T) {
	// Call GetCaller from nested anonymous functions
	caller := func() commonz.CallerInfo {
		return func() commonz.CallerInfo {
			return func() commonz.CallerInfo {
				return commonz.GetCaller(commonz.GrandparentCaller)
			}()
		}()
	}()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "TestGetCallerFromNestedAnonymousFunctions.func1",
	}

	require.Equal(t, expected, caller, "GetCaller(GrandparentCaller) from nested anonymous functions should return correct caller info")
}

// TestGetCallerFromAnonymousFunctionInMethod tests GetCaller when called from an anonymous function inside a method
func TestGetCallerFromAnonymousFunctionInMethod(t *testing.T) {
	// Create a test struct with a method that contains an anonymous function
	ts := &TestStruct{}
	caller := ts.getCallerFromAnonymous()

	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "TestGetCallerFromAnonymousFunctionInMethod.(*TestStruct).getCallerFromAnonymous.func1",
	}

	require.Equal(t, expected, caller, "GetCaller(CurrentCaller) from anonymous function inside method should return correct caller info")
}

// TestInitCaller tests that the package-level initCaller variable has the expected values
func TestInitCaller(t *testing.T) {
	// The initCaller variable should capture the package initialization context
	// It should have the correct package name and a function name that indicates package initialization
	expected := commonz.CallerInfo{
		Package:  "github.com/goosz/commonz_test",
		Function: "init", // Package initialization function
	}

	require.Equal(t, expected, initCaller, "initCaller should capture the package initialization context")
}

// Simple exported function for testing GetCaller
func ExportedTestFunction() commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}

// Simple unexported function for testing GetCaller
func unexportedTestFunction() commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}

// Generic simple function for testing GetCaller
func GenericTestFunction[T any](value T) commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}

// Test structs for receiver method testing
type ValueReceiverStruct struct {
	name string
}

// Value receiver method that calls GetCaller
func (v ValueReceiverStruct) GetCallerInfo() commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}

type PointerReceiverStruct struct {
	name string
}

// Pointer receiver method that calls GetCaller
func (p *PointerReceiverStruct) GetCallerInfo() commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}

// Generic receiver structs for testing
type GenericValueReceiver[T any] struct {
	value T
}

// Value receiver method on generic type that calls GetCaller
func (g GenericValueReceiver[T]) GetCallerInfo() commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}

type GenericPointerReceiver[T any] struct {
	value T
}

// Pointer receiver method on generic type that calls GetCaller
func (g *GenericPointerReceiver[T]) GetCallerInfo() commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}

// Embedded structs for testing
type EmbeddedValueStruct struct {
	ValueReceiverStruct // embedded value receiver struct
	extraValue          int
}

type EmbeddedPointerStruct struct {
	*PointerReceiverStruct // embedded pointer receiver struct
	extraValue             int
}

// TestStruct for testing anonymous functions in methods
type TestStruct struct{}

// getCallerFromAnonymous calls GetCaller from an anonymous function
func (ts *TestStruct) getCallerFromAnonymous() commonz.CallerInfo {
	return func() commonz.CallerInfo {
		return commonz.GetCaller(commonz.CurrentCaller)
	}()
}
