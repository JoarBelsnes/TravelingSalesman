package salesman

import "testing"

func benchmarkSalesman(i int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < i; n++ {
		for n := 0; n < b.N; n++ {
			Salesman()
		}
	}
}

func BenchmarkSalesman1(b *testing.B)   { benchmarkSalesman(1, b) }
func BenchmarkSalesman10(b *testing.B)  { benchmarkSalesman(10, b) }
func BenchmarkSalesman100(b *testing.B) { benchmarkSalesman(100, b) }
