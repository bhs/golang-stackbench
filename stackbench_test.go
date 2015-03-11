package stackbench

import (
	"runtime"
	"testing"
)

// A helper that generates a stack trace of depth >= n, then writes the
// Stack()-trace into a buffer of max size traceBufferSize.
func recurseNThenCallClosure(n int, closure func(), readyDoneSignal chan<- bool) {
	if n == 0 {
		readyDoneSignal <- true // ready
		closure()
		readyDoneSignal <- true // done
	} else {
		recurseNThenCallClosure(n-1, closure, readyDoneSignal)
	}
}

/////////////////////////
// LENGTH=10K STACKTRACES
/////////////////////////

// Helper for benchmarking runtime.Stack().
func benchmarkStacktrace(b *testing.B, stackDepth, traceBufferSize int) {
	// recurseNThenCallClosure sends to readyDoneSignal once when it's ready, then once more when it's done.
	readyDoneSignal := make(chan bool)
	go recurseNThenCallClosure(stackDepth, func() {
		for i := 0; i < b.N; i++ {
			// Take our stack trace.
			buffer := make([]byte, traceBufferSize)
			_ = runtime.Stack(buffer, false)
		}
	}, readyDoneSignal)
	<-readyDoneSignal // ready
	b.ResetTimer()
	<-readyDoneSignal // done
}

// Helper for benchmarking runtime.Callers().
func benchmarkCallers(b *testing.B, stackDepth, maxStackDepth int) {
	// recurseNThenCallClosure sends to readyDoneSignal once when it's ready, then once more when it's done.
	readyDoneSignal := make(chan bool)
	go recurseNThenCallClosure(stackDepth, func() {
		for i := 0; i < b.N; i++ {
			// Take our stack trace.
			buffer := make([]uintptr, maxStackDepth)
			_ = runtime.Callers(1, buffer)
		}
	}, readyDoneSignal)
	<-readyDoneSignal // ready
	b.ResetTimer()
	<-readyDoneSignal // done
}

func BenchmarkStacktracesDepth5Length10K(b *testing.B) {
	benchmarkStacktrace(b, 5, 10000)
}

func BenchmarkStacktracesDepth10Length10K(b *testing.B) {
	benchmarkStacktrace(b, 10, 10000)
}

func BenchmarkStacktracesDepth50Length10K(b *testing.B) {
	benchmarkStacktrace(b, 50, 10000)
}

func BenchmarkStacktracesDepth100Length10K(b *testing.B) {
	benchmarkStacktrace(b, 100, 10000)
}

////////////////////////
// LENGTH=1K STACKTRACES
////////////////////////

func BenchmarkStacktracesDepth5Length1K(b *testing.B) {
	benchmarkStacktrace(b, 5, 1000)
}

func BenchmarkStacktracesDepth10Length1K(b *testing.B) {
	benchmarkStacktrace(b, 10, 1000)
}

func BenchmarkStacktracesDepth50Length1K(b *testing.B) {
	benchmarkStacktrace(b, 50, 1000)
}

func BenchmarkStacktracesDepth100Length1K(b *testing.B) {
	benchmarkStacktrace(b, 100, 1000)
}

/////////////////////////
// LENGTH=100 STACKTRACES
/////////////////////////

func BenchmarkStacktracesDepth5Length100(b *testing.B) {
	benchmarkStacktrace(b, 5, 100)
}

func BenchmarkStacktracesDepth10Length100(b *testing.B) {
	benchmarkStacktrace(b, 10, 100)
}

func BenchmarkStacktracesDepth50Length100(b *testing.B) {
	benchmarkStacktrace(b, 50, 100)
}

func BenchmarkStacktracesDepth100Length100(b *testing.B) {
	benchmarkStacktrace(b, 100, 100)
}

//////////////////////
// MAX_DEPTH=5 CALLERS
//////////////////////

func BenchmarkCallersDepth5Limit5(b *testing.B) {
	benchmarkCallers(b, 5, 5)
}

func BenchmarkCallersDepth10Limit5(b *testing.B) {
	benchmarkCallers(b, 10, 5)
}

func BenchmarkCallersDepth50Limit5(b *testing.B) {
	benchmarkCallers(b, 50, 5)
}

func BenchmarkCallersDepth100Limit5(b *testing.B) {
	benchmarkCallers(b, 100, 5)
}

///////////////////////
// MAX_DEPTH=50 CALLERS
///////////////////////

func BenchmarkCallersDepth5Limit100(b *testing.B) {
	benchmarkCallers(b, 5, 100)
}

func BenchmarkCallersDepth10Limit100(b *testing.B) {
	benchmarkCallers(b, 10, 100)
}

func BenchmarkCallersDepth50Limit100(b *testing.B) {
	benchmarkCallers(b, 50, 100)
}

func BenchmarkCallersDepth100Limit100(b *testing.B) {
	benchmarkCallers(b, 100, 100)
}

// A test that does nothing more than silence the `go test` warning about a
// lack of tests. :)
func TestNothing(t *testing.T) {}
