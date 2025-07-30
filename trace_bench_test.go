package commonz_test

import (
	"testing"

	"github.com/goosz/commonz"
)

// BenchmarkGetCaller benchmarks the GetCaller function at different depths
func BenchmarkGetCaller(b *testing.B) {
	b.Run("depth_0", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = commonz.GetCaller(commonz.CurrentCaller)
		}
	})

	b.Run("depth_1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = commonz.GetCaller(commonz.ParentCaller)
		}
	})

	b.Run("depth_5", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = commonz.GetCaller(5) // Using hardcoded value for benchmarking specific depth
		}
	})

	b.Run("depth_10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = commonz.GetCaller(10) // Using hardcoded value for benchmarking specific depth
		}
	})

	b.Run("negative_depth", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = commonz.GetCaller(-1)
		}
	})

	b.Run("large_depth", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = commonz.GetCaller(100) // Using hardcoded value for benchmarking edge case
		}
	})
}

// BenchmarkGetCaller_fromMethod benchmarks GetCaller when called from a method
func BenchmarkGetCaller_fromMethod(b *testing.B) {
	processor := &BenchmarkDataProcessor{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processor.GetCallerInfo()
	}
}

// BenchmarkGetCaller_fromAnonymousFunction benchmarks GetCaller from anonymous functions
func BenchmarkGetCaller_fromAnonymousFunction(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = func() commonz.CallerInfo {
			return commonz.GetCaller(commonz.CurrentCaller)
		}()
	}
}

// BenchmarkParseCallerInfo benchmarks ParseCallerInfo with different function name types
func BenchmarkParseCallerInfo(b *testing.B) {
	testCases := []struct {
		name   string
		fnName string
	}{
		{"simple_function", "github.com/user/package.SimpleFunction"},
		{"value_method", "github.com/user/package.MyStruct.GetValue"},
		{"pointer_method", "github.com/user/package.(*MyStruct).SetValue"},
		{"generic_function", "github.com/user/package.GenericFunction[string]"},
		{"generic_method", "github.com/user/package.(*GenericType[int]).GetValue"},
		{"nested_function", "github.com/user/package.OuterFunction.InnerFunction"},
		{"complex_generic", "github.com/user/package.Container[map[string]interface{}].Process"},
		{"empty_string", ""},
		{"no_periods", "justafunctionname"},
		{"malformed", "package.(*Type.GetMethod"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = commonz.ParseCallerInfo(tc.fnName)
			}
		})
	}
}

// BenchmarkCallerInfo_String benchmarks the String method
func BenchmarkCallerInfo_String(b *testing.B) {
	testCases := []struct {
		name string
		info commonz.CallerInfo
	}{
		{"simple", commonz.CallerInfo{Package: "github.com/user/package", Function: "SimpleFunction"}},
		{"method", commonz.CallerInfo{Package: "github.com/user/package", Function: "MyStruct.GetValue"}},
		{"pointer_method", commonz.CallerInfo{Package: "github.com/user/package", Function: "(*MyStruct).SetValue"}},
		{"generic", commonz.CallerInfo{Package: "github.com/user/package", Function: "GenericFunction[string]"}},
		{"complex", commonz.CallerInfo{Package: "github.com/user/package", Function: "(*GenericType[int]).GetValue"}},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = tc.info.String()
			}
		})
	}
}

// BenchmarkCallerInfo_IsUnknown benchmarks the IsUnknown method
func BenchmarkCallerInfo_IsUnknown(b *testing.B) {
	testCases := []struct {
		name string
		info commonz.CallerInfo
	}{
		{"known", commonz.CallerInfo{Package: "github.com/user/package", Function: "SimpleFunction"}},
		{"unknown_package", commonz.CallerInfo{Package: "<unknown-package>", Function: "SimpleFunction"}},
		{"unknown_function", commonz.CallerInfo{Package: "github.com/user/package", Function: "<unknown-function>"}},
		{"both_unknown", commonz.CallerInfo{Package: "<unknown-package>", Function: "<unknown-function>"}},
		{"empty", commonz.CallerInfo{Package: "", Function: ""}},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = tc.info.IsUnknown()
			}
		})
	}
}

// BenchmarkParseCallerInfo_vs_GetCaller benchmarks the relative performance
func BenchmarkParseCallerInfo_vs_GetCaller(b *testing.B) {
	// Pre-computed function names to avoid the overhead of runtime.Caller
	functionNames := []string{
		"github.com/user/package.SimpleFunction",
		"github.com/user/package.MyStruct.GetValue",
		"github.com/user/package.(*MyStruct).SetValue",
		"github.com/user/package.GenericFunction[string]",
		"github.com/user/package.(*GenericType[int]).GetValue",
	}

	b.Run("ParseCallerInfo_only", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fnName := functionNames[i%len(functionNames)]
			_ = commonz.ParseCallerInfo(fnName)
		}
	})

	b.Run("GetCaller_only", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = commonz.GetCaller(commonz.CurrentCaller)
		}
	})

	b.Run("GetCaller_with_parsing", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			caller := commonz.GetCaller(commonz.CurrentCaller)
			_ = caller.String()
		}
	})
}

// BenchmarkConcurrentGetCaller benchmarks GetCaller under concurrent access
func BenchmarkConcurrentGetCaller(b *testing.B) {
	b.Run("concurrent_depth_0", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = commonz.GetCaller(commonz.CurrentCaller)
			}
		})
	})

	b.Run("concurrent_depth_1", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_ = commonz.GetCaller(commonz.ParentCaller)
			}
		})
	})
}

// BenchmarkConcurrentParseCallerInfo benchmarks ParseCallerInfo under concurrent access
func BenchmarkConcurrentParseCallerInfo(b *testing.B) {
	functionNames := []string{
		"github.com/user/package.SimpleFunction",
		"github.com/user/package.MyStruct.GetValue",
		"github.com/user/package.(*MyStruct).SetValue",
		"github.com/user/package.GenericFunction[string]",
		"github.com/user/package.(*GenericType[int]).GetValue",
	}

	b.Run("concurrent_simple", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			i := 0
			for pb.Next() {
				fnName := functionNames[i%len(functionNames)]
				_ = commonz.ParseCallerInfo(fnName)
				i++
			}
		})
	})
}

// Helper types for benchmarks

// BenchmarkDataProcessor is a helper struct for method benchmarks
type BenchmarkDataProcessor struct{}

// GetCallerInfo returns caller info from a method
func (dp *BenchmarkDataProcessor) GetCallerInfo() commonz.CallerInfo {
	return commonz.GetCaller(commonz.CurrentCaller)
}
