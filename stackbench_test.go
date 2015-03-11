package stackbench

import (
	"runtime"
	"testing"
)

// A helper that generates a stack trace of depth >= n, then writes the
// Stack()-trace into a buffer of max size traceBufferSize.
func recurseNThenStackTrace(n int, traceBufferSize int, stacktraceSignal <-chan int, readyDoneSignal chan<- bool) {
	if n == 0 {
		readyDoneSignal <- true // ready
		count := <-stacktraceSignal
		for i := 0; i < count; i++ {
			// Take our stack trace.
			buffer := make([]byte, traceBufferSize)
			_ = runtime.Stack(buffer, false)
		}
		readyDoneSignal <- true // done
	} else {
		recurseNThenStackTrace(n-1, traceBufferSize, stacktraceSignal, readyDoneSignal)
	}
}

/////////////////////////
// LENGTH=10K STACKTRACES
/////////////////////////

func benchmarkStacktrace(b *testing.B, stackDepth, traceBufferSize int) {
	stacktraceSignal := make(chan int)
	// recurseNThenStackTrace sends to readyDoneSignal once when it's ready, then once more when it's done.
	readyDoneSignal := make(chan bool)
	go recurseNThenStackTrace(stackDepth, traceBufferSize, stacktraceSignal, readyDoneSignal)
	<-readyDoneSignal
	b.ResetTimer()
	stacktraceSignal <- b.N
	<-readyDoneSignal
}

func BenchmarkDepth5Length10KStacktraces(b *testing.B) {
	benchmarkStacktrace(b, 5, 10000)
}

func BenchmarkDepth10Length10KStacktraces(b *testing.B) {
	benchmarkStacktrace(b, 10, 10000)
}

func BenchmarkDepth50Length10KStacktraces(b *testing.B) {
	benchmarkStacktrace(b, 50, 10000)
}

func BenchmarkDepth100Length10KStacktraces(b *testing.B) {
	benchmarkStacktrace(b, 100, 10000)
}

////////////////////////
// LENGTH=1K STACKTRACES
////////////////////////

func BenchmarkDepth5Length1KStacktraces(b *testing.B) {
	benchmarkStacktrace(b, 5, 1000)
}

func BenchmarkDepth10Length1KStacktraces(b *testing.B) {
	benchmarkStacktrace(b, 10, 1000)
}

func BenchmarkDepth50Length1KStacktraces(b *testing.B) {
	benchmarkStacktrace(b, 50, 1000)
}

func BenchmarkDepth100Length1KStacktraces(b *testing.B) {
	benchmarkStacktrace(b, 100, 1000)
}

/////////////////////////
// LENGTH=100 STACKTRACES
/////////////////////////

func BenchmarkDepth5Length100Stacktraces(b *testing.B) {
	benchmarkStacktrace(b, 5, 100)
}

func BenchmarkDepth10Length100Stacktraces(b *testing.B) {
	benchmarkStacktrace(b, 10, 100)
}

func BenchmarkDepth50Length100Stacktraces(b *testing.B) {
	benchmarkStacktrace(b, 50, 100)
}

func BenchmarkDepth100Length100Stacktraces(b *testing.B) {
	benchmarkStacktrace(b, 100, 100)
}

// A test that does nothing more than silence the `go test` warning about a
// lack of tests. :)
func TestNothing(t *testing.T) {}
