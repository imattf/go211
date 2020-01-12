package word

import (
	"fmt"
	"testing"

	"github.com/imattf/go211/ninj1300/02/quote"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "expected", 1)
			}
		case "three":
			if v != 2 {
				t.Error("got", v, "expected", 3)
			}
		}
	}
}
func TestCount(t *testing.T) {

	type test struct {
		data   string
		answer int
	}
	tests := []test{
		test{"1 2", 2},
		// test{3, 21},
		// test{5, 35},
	}

	for _, v := range tests {
		x := Count(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}

	}
}

func ExampleCount() {
	fmt.Println(Count("1 2"))
	// Output:
	// 2
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}
