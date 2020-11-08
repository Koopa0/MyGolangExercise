package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "Shaken not stirred"{
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred"{
		t.Error("got", s, "want", "Shaken not stirred")
	}
}

func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))

}

const s = "Hello"

var xs [] string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i<b.N;i++ {
		Cat(xs)
	}
}
func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0;i <b.N; i++ {
		Join(xs)

	}
}