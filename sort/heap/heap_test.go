package heap

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	t.Run("Testing heap sort with simple array", func(t *testing.T) {
		x := []int{8, 5, 3, 10, 36}
		expected := []int{3, 5, 8, 10, 36}
		HeapSort(x)
		for i, v := range x {
			if v != expected[i] {
				t.Error("Test failed for heap sort with simple array at the index", i)
			}
		}
	})
}

func BenchmarkHeapSort(b *testing.B) {
	b.Run("Benchmarking heap sort with simple array", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x := []int{8, 5, 3, 10, 36}
			HeapSort(x)
		}
	})
}

func ExampleHeapSort() {
	x := []int{8, 5, 3, 10, 36}
	HeapSort(x)
	fmt.Println(x)
	//Expected output: [3 5 8 10 36]
}
