package stackbench

import (
	"runtime"
	"testing"
)

// A helper that generates a stack trace of depth >= n, then writes the
// Stack()-trace into a buffer of max size traceBufferSize.
func recurseNThenStackTrace(n int, traceBufferSize int) {
	if n == 0 {
		// Take our stack trace.
		buffer := make([]byte, traceBufferSize)
		_ = runtime.Stack(buffer, false)
	} else {
		recurseNThenStackTrace(n-1, traceBufferSize)
	}
}

/////////////////////////
// LENGTH=10K STACKTRACES
/////////////////////////

func BenchmarkDepth5Length10KStacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(5, 10000)
	}
}

func BenchmarkDepth10Length10KStacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(10, 10000)
	}
}

func BenchmarkDepth50Length10KStacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(50, 10000)
	}
}

func BenchmarkDepth100Length10KStacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(100, 10000)
	}
}

////////////////////////
// LENGTH=1K STACKTRACES
////////////////////////

func BenchmarkDepth5Length1KStacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(5, 1000)
	}
}

func BenchmarkDepth10Length1KStacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(10, 1000)
	}
}

func BenchmarkDepth50Length1KStacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(50, 1000)
	}
}

func BenchmarkDepth100Length1KStacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(100, 1000)
	}
}

/////////////////////////
// LENGTH=100 STACKTRACES
/////////////////////////

func BenchmarkDepth5Length100Stacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(5, 100)
	}
}

func BenchmarkDepth10Length100Stacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(10, 100)
	}
}

func BenchmarkDepth50Length100Stacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(50, 100)
	}
}

func BenchmarkDepth100Length100Stacktraces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		recurseNThenStackTrace(100, 100)
	}
}

// A test that does nothing more than silence the `go test` warning about a
// lack of tests. :)
func TestNothing(t *testing.T) {}
