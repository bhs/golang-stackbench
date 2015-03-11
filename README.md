# golang-stack-bench

Benchmarking the golang runtime.Stack(...) function in terms of the stack depth
and output buffer size.

Here are the results I get on my 2.3GHz MBP for ballparking purposes:

    $ go test -bench . stackbench
    PASS
    BenchmarkStacktracesDepth5Length10K	  100000	     22264 ns/op
    BenchmarkStacktracesDepth10Length10K	   50000	     27062 ns/op
    BenchmarkStacktracesDepth50Length10K	   20000	     63515 ns/op
    BenchmarkStacktracesDepth100Length10K	   10000	    107113 ns/op
    BenchmarkStacktracesDepth5Length1K	  200000	      8973 ns/op
    BenchmarkStacktracesDepth10Length1K	  100000	     13399 ns/op
    BenchmarkStacktracesDepth50Length1K	   30000	     47361 ns/op
    BenchmarkStacktracesDepth100Length1K	   20000	     88230 ns/op
    BenchmarkStacktracesDepth5Length100	  300000	      6799 ns/op
    BenchmarkStacktracesDepth10Length100	  200000	     10950 ns/op
    BenchmarkStacktracesDepth50Length100	   30000	     44386 ns/op
    BenchmarkStacktracesDepth100Length100	   20000	     84219 ns/op
    BenchmarkCallersDepth5Limit5	 2000000	       953 ns/op
    BenchmarkCallersDepth10Limit5	 2000000	       956 ns/op
    BenchmarkCallersDepth50Limit5	 2000000	       972 ns/op
    BenchmarkCallersDepth100Limit5	 2000000	       988 ns/op
    BenchmarkCallersDepth5Limit100	 1000000	      3245 ns/op
    BenchmarkCallersDepth10Limit100	  500000	      3925 ns/op
    BenchmarkCallersDepth50Limit100	  200000	      9365 ns/op
    BenchmarkCallersDepth100Limit100	  100000	     15691 ns/op
    ok  	stackbench	44.695s

