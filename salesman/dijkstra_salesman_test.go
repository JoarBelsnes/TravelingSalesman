package salesman

import "testing"

func TestSalesman(t *testing.T) {
	Salesman()
}

/*
func benchmarkSalesman(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Salesman()

	}
}

func BenchmarkSalesman1(b *testing.B) { benchmarkSalesman(b) }
*/
/*func BenchmarkSalesman10(b *testing.B)  { benchmarkSalesman(10, b) }
func BenchmarkSalesman100(b *testing.B) { benchmarkSalesman(100, b) }
*/
