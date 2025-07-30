package commonz_test

import (
	"fmt"
	"log"

	"github.com/goosz/commonz"
)

// ExampleGetCaller demonstrates how to get caller information at different depths
func ExampleGetCaller() {
	// Get information about the current function
	current := commonz.GetCaller(commonz.CurrentCaller)
	fmt.Printf("Current function: %s.%s\n", current.Package, current.Function)

	// Get information about the caller of this function
	caller := commonz.GetCaller(commonz.ParentCaller)
	fmt.Printf("Caller function: %s.%s\n", caller.Package, caller.Function)

	// Output:
	// Current function: github.com/goosz/commonz_test.ExampleGetCaller
	// Caller function: testing.runExample
}

// ExampleGetCaller_fromMethod demonstrates getting caller info from a method
func ExampleGetCaller_fromMethod() {
	// Create a struct and call a method
	processor := &DataProcessor{name: "test-processor"}
	info := processor.ProcessData("sample data")

	fmt.Printf("Called from: %s.%s\n", info.Package, info.Function)

	// Output:
	// Called from: github.com/goosz/commonz_test.(*DataProcessor).ProcessData
}

// ExampleGetCaller_fromAnonymousFunction demonstrates getting caller info from an anonymous function
func ExampleGetCaller_fromAnonymousFunction() {
	// Call GetCaller from an anonymous function
	info := func() commonz.CallerInfo {
		return commonz.GetCaller(commonz.CurrentCaller)
	}()

	fmt.Printf("Anonymous function: %s.%s\n", info.Package, info.Function)

	// Output:
	// Anonymous function: github.com/goosz/commonz_test.ExampleGetCaller_fromAnonymousFunction.func1
}

// ExampleGetCaller_withDepth demonstrates using different depth values
func ExampleGetCaller_withDepth() {
	// Call through multiple levels
	info := callThroughMultipleLevels()

	fmt.Printf("Original caller: %s.%s\n", info.Package, info.Function)

	// Output:
	// Original caller: github.com/goosz/commonz_test.callThroughMultipleLevels
}

// ExampleParseCallerInfo demonstrates parsing function names
func ExampleParseCallerInfo() {
	// Parse a simple function name
	simple := commonz.ParseCallerInfo("github.com/user/package.SimpleFunction")
	fmt.Printf("Simple function: %s.%s\n", simple.Package, simple.Function)

	// Parse a value receiver method
	valueMethod := commonz.ParseCallerInfo("github.com/user/package.MyStruct.GetValue")
	fmt.Printf("Value method: %s.%s\n", valueMethod.Package, valueMethod.Function)

	// Parse a pointer receiver method
	pointerMethod := commonz.ParseCallerInfo("github.com/user/package.(*MyStruct).SetValue")
	fmt.Printf("Pointer method: %s.%s\n", pointerMethod.Package, pointerMethod.Function)

	// Parse a generic function
	generic := commonz.ParseCallerInfo("github.com/user/package.GenericFunction[string]")
	fmt.Printf("Generic function: %s.%s\n", generic.Package, generic.Function)

	// Output:
	// Simple function: github.com/user/package.SimpleFunction
	// Value method: github.com/user/package.MyStruct.GetValue
	// Pointer method: github.com/user/package.(*MyStruct).SetValue
	// Generic function: github.com/user/package.GenericFunction[string]
}

// ExampleParseCallerInfo_errorCases demonstrates handling error cases
func ExampleParseCallerInfo_errorCases() {
	// Parse an empty string
	empty := commonz.ParseCallerInfo("")
	fmt.Printf("Empty string: %s.%s\n", empty.Package, empty.Function)
	fmt.Printf("Is unknown: %t\n", empty.IsUnknown())

	// Parse a string with no periods
	noPeriods := commonz.ParseCallerInfo("justafunctionname")
	fmt.Printf("No periods: %s.%s\n", noPeriods.Package, noPeriods.Function)
	fmt.Printf("Is unknown: %t\n", noPeriods.IsUnknown())

	// Parse a malformed function name
	malformed := commonz.ParseCallerInfo("package.(*Type.GetMethod")
	fmt.Printf("Malformed: %s.%s\n", malformed.Package, malformed.Function)
	fmt.Printf("Is unknown: %t\n", malformed.IsUnknown())

	// Output:
	// Empty string: <unknown-package>.<unknown-function>
	// Is unknown: true
	// No periods: <unknown-package>.<unknown-function>
	// Is unknown: true
	// Malformed: package.(*Type.GetMethod
	// Is unknown: false
}

// ExampleCallerInfo_String demonstrates the String method
func ExampleCallerInfo_String() {
	// Create a CallerInfo and get its string representation
	info := commonz.CallerInfo{
		Package:  "github.com/user/package",
		Function: "MyStruct.GetValue",
	}

	fmt.Println("String representation:", info.String())

	// Output:
	// String representation: github.com/user/package.MyStruct.GetValue
}

// ExampleCallerInfo_IsUnknown demonstrates the IsUnknown method
func ExampleCallerInfo_IsUnknown() {
	// Create a known caller info
	known := commonz.CallerInfo{
		Package:  "github.com/user/package",
		Function: "MyFunction",
	}
	fmt.Printf("Known caller is unknown: %t\n", known.IsUnknown())

	// Create an unknown caller info
	unknown := commonz.CallerInfo{
		Package:  "<unknown-package>",
		Function: "<unknown-function>",
	}
	fmt.Printf("Unknown caller is unknown: %t\n", unknown.IsUnknown())

	// Output:
	// Known caller is unknown: false
	// Unknown caller is unknown: true
}

// ExampleGetCaller_logging demonstrates using GetCaller for logging
func ExampleGetCaller_logging() {
	// Simulate a logging scenario
	logger := &SimpleLogger{}

	logger.Info("Processing started")
	logger.Debug("Data validation passed")
	logger.Error("Failed to connect to database")

	// Output:
	// [INFO] github.com/goosz/commonz_test.ExampleGetCaller_logging: Processing started
	// [DEBUG] github.com/goosz/commonz_test.ExampleGetCaller_logging: Data validation passed
	// [ERROR] github.com/goosz/commonz_test.ExampleGetCaller_logging: Failed to connect to database
}

// ExampleGetCaller_debugging demonstrates using GetCaller for debugging
func ExampleGetCaller_debugging() {
	// Simulate a debugging scenario
	debugger := &SimpleDebugger{}

	debugger.Trace("Entering function")
	debugger.Trace("Processing data")
	debugger.Trace("Exiting function")

	// Output:
	// [TRACE] github.com/goosz/commonz_test.ExampleGetCaller_debugging: Entering function
	// [TRACE] github.com/goosz/commonz_test.ExampleGetCaller_debugging: Processing data
	// [TRACE] github.com/goosz/commonz_test.ExampleGetCaller_debugging: Exiting function
}

// Helper functions and types for examples

// DataProcessor is a sample struct for demonstrating method calls
type DataProcessor struct {
	name string
}

// ProcessData demonstrates getting caller info from a method
func (dp *DataProcessor) ProcessData(data string) commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}

// callThroughMultipleLevels demonstrates call stack depth
func callThroughMultipleLevels() commonz.CallerInfo {
	return level3()
}

func level3() commonz.CallerInfo {
	return level2()
}

func level2() commonz.CallerInfo {
	return level1()
}

func level1() commonz.CallerInfo {
	return commonz.GetCaller(3) // Go back 3 levels to the original caller (using hardcoded value for demonstration)
}

// Logger demonstrates using GetCaller for logging
type Logger struct{}

func (l *Logger) Info(message string) {
	caller := commonz.GetCaller(commonz.CurrentCaller)
	log.Printf("[INFO] %s: %s", caller.String(), message)
}

func (l *Logger) Debug(message string) {
	caller := commonz.GetCaller(commonz.CurrentCaller)
	log.Printf("[DEBUG] %s: %s", caller.String(), message)
}

func (l *Logger) Error(message string) {
	caller := commonz.GetCaller(commonz.CurrentCaller)
	log.Printf("[ERROR] %s: %s", caller.String(), message)
}

// Debugger demonstrates using GetCaller for debugging
type Debugger struct{}

func (d *Debugger) Trace(message string) {
	caller := commonz.GetCaller(commonz.CurrentCaller)
	log.Printf("[TRACE] %s: %s", caller.String(), message)
}

// SimpleLogger demonstrates using GetCaller for logging without timestamps
type SimpleLogger struct{}

func (l *SimpleLogger) Info(message string) {
	caller := commonz.GetCaller(commonz.ParentCaller) // Get the caller of this method
	fmt.Printf("[INFO] %s: %s\n", caller.String(), message)
}

func (l *SimpleLogger) Debug(message string) {
	caller := commonz.GetCaller(commonz.ParentCaller) // Get the caller of this method
	fmt.Printf("[DEBUG] %s: %s\n", caller.String(), message)
}

func (l *SimpleLogger) Error(message string) {
	caller := commonz.GetCaller(commonz.ParentCaller) // Get the caller of this method
	fmt.Printf("[ERROR] %s: %s\n", caller.String(), message)
}

// SimpleDebugger demonstrates using GetCaller for debugging without timestamps
type SimpleDebugger struct{}

func (d *SimpleDebugger) Trace(message string) {
	caller := commonz.GetCaller(commonz.ParentCaller) // Get the caller of this method
	fmt.Printf("[TRACE] %s: %s\n", caller.String(), message)
}
