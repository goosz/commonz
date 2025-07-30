package commonz

import (
	"fmt"
	"runtime"
	"strings"
)

// Common depth constants for GetCaller function
const (
	CurrentCaller     = 0 // The function that called GetCaller
	ParentCaller      = 1 // The function that called the function that called GetCaller
	GrandparentCaller = 2 // The function that called the function that called the function that called GetCaller
)

// CallerInfo contains the parsed components of a Go function name.
type CallerInfo struct {
	Package  string // The fully qualified package name (e.g., "github.com/goosz/commonz_test")
	Function string // Everything after the package name, including function name, type
	// information for methods, etc. (e.g., "TestFunction", "MyType.Method", "(*MyType).Method")
}

func (ci CallerInfo) String() string {
	return fmt.Sprintf("%s.%s", ci.Package, ci.Function)
}

// IsUnknown returns true if the CallerInfo represents an unknown caller.
// This is the case when the package is "<unknown-package>" or the function is "<unknown-function>".
func (ci CallerInfo) IsUnknown() bool {
	return ci.Package == "<unknown-package>" || ci.Function == "<unknown-function>"
}

// ParseCallerInfo extracts the components of a Go function name.
//
// Function names can be in various formats:
// - "package.function" - simple function
// - "package.type.method" - method on a value type
// - "package.(*type).method" - method on a pointer type
// - "full/import/path.function" - imported package function
// - "full/import/path.(*type).method" - imported package method
// - "full/import/path.init" - init function
//
// Returns a CallerInfo struct with Package and Function components.
// The Function field contains everything after the package name, preserving
// type information, method names, and any nested function structure.
// If the function name cannot be parsed, returns a CallerInfo with Package set to "<unknown-package>".
func ParseCallerInfo(fnName string) CallerInfo {
	// Find the first period starting after the last slash (or from beginning if no slashes)
	startPos := strings.LastIndexByte(fnName, '/')
	if startPos < 0 {
		startPos = 0
	} else {
		startPos++
	}
	if periodPos := strings.IndexByte(fnName[startPos:], '.'); periodPos >= 0 {
		return CallerInfo{
			Package:  fnName[:startPos+periodPos],
			Function: fnName[startPos+periodPos+1:],
		}
	}

	return unknownCallerInfo()
}

// unknownCallerInfo returns a CallerInfo struct representing an unknown caller
func unknownCallerInfo() CallerInfo {
	return CallerInfo{
		Package:  "<unknown-package>",
		Function: "<unknown-function>",
	}
}

// GetCaller returns the CallerInfo at the specified depth in the call stack.
//
// This function uses runtime introspection to determine the caller information
// at the given call stack depth. It's useful for adding detailed caller context
// to debug information, logging, or identity systems.
//
// The depth parameter specifies how many levels up the call stack to look.
// Common depth values are provided as constants:
// - CurrentCaller (0): the function that called GetCaller
// - ParentCaller (1): the function that called the function that called GetCaller
// - GrandparentCaller (2): the function that called the function that called the function that called GetCaller
// - Higher values: continue up the call stack as needed
//
// Returns a CallerInfo struct with Package and Function components.
// If the caller cannot be determined at the specified depth, returns a CallerInfo
// with Package set to "<unknown-package>" and Function set to "<unknown-function>".
// Use the IsUnknown() method to check if the returned CallerInfo represents an unknown caller.
func GetCaller(depth int) CallerInfo {
	if depth < 0 {
		return unknownCallerInfo()
	}

	// Get the caller at the specified depth in the call stack
	pc, _, _, ok := runtime.Caller(depth + 1) // +1 because Caller(0) would be GetCaller itself
	if !ok {
		return unknownCallerInfo()
	}

	// Get the function information from the program counter
	if fn := runtime.FuncForPC(pc); fn != nil {
		return ParseCallerInfo(fn.Name())
	}

	return unknownCallerInfo()
}
