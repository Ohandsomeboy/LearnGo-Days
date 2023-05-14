package start

import "testing"

// 基准测试默认执行实践是1秒

// BenchmarkFib10 run 'go test -bench=. -benchmem' to get the benchmark result
func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
