package iteration

import "testing"
import "fmt"

const repeatCount = 5

func Repeat (character string, iters ...int) string {
	var repeated string 
	for i := 0; i < iters[0]; i ++ {
		repeated += character
	}
	return repeated
}

func BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a",100)
	}
}

func TestRepeat(t *testing.T){
	repeated := Repeat("a",10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat(){
	repeated := Repeat("a",10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}