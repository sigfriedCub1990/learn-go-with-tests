package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
    sum := Add(2, 2)
    expected := 4

    if sum != expected {
        t.Errorf("Expected '%d' but go '%d'", expected, sum)
    }
}

func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
