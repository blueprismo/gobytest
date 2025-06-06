package integers

import "testing"
import "fmt"

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}

func TestAdder(t *testing.T){
	sum := Add(2,2)
	expected := 4
	if (sum != expected){
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Examples
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}