package dog

import (
	"fmt"
	"testing"
)

//Years stuff...
func TestYears(t *testing.T) {

	type test struct {
		data   int
		answer int
	}
	tests := []test{
		test{0, 0},
		test{3, 21},
		test{5, 35},
	}

	for _, v := range tests {
		x := Years(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}

	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Output:
	// 21
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}
}

//YearsTwo stuff...
func TestYearsTwo(t *testing.T) {

	type test struct {
		data   int
		answer int
	}
	tests := []test{
		test{0, 0},
		test{3, 21},
		test{5, 35},
	}

	for _, v := range tests {
		x := YearsTwo(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}

	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	// Output:
	// 21
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}
}
