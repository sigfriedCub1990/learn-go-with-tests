package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
    t.Run("repeats N times", func (t *testing.T) {
        repeated := Repeat("a", 6)
        expected := "aaaaaa"

        if repeated != expected {
            t.Errorf("Expected '%q' but got '%q'", expected, repeated)
        }
    })

}

func ExampleRepeat() {
    repeated := Repeat("a", 5)
    fmt.Println(repeated)
    // Output: aaaaa
}

// For our implementation this yields: 90.53 ns/op
// For Go's implementation with strings.Repeat() this yields 38.70 ns/op!
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}
