package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Koopa")
	if s != "Welcome my dear Koopa"{
		t.Error("got", s, "expected", "Welcome my dear Koopa")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Koopa"))
}
func BenchmarkGreet(b *testing.B) {
	for i := 0;i < b.N; i++{
		Greet("Koopa")
	}
}