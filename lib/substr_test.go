package lib

import (
	"fmt"
	"testing"
)

var s = "Google事"


func TestSubStrRune(t *testing.T) {
	res := SubStrRune(s, 1, 4)
	if res != "oog" {
		t.Errorf("%s is not fit", res)
	}
}

func BenchmarkSubStrRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubStrRune(s, 1, 4)
	}
}

func ExampleSubStrRune() {
	res := SubStrRune(s, 0, 4)
	fmt.Println(res)
	// Output: Goog
}