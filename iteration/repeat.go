package iteration

import "strings"

func Repeat(character string, times int) string {
    // NOTE: We could have used strings.Repeat(character, times)
    // var repeated string
    // for i := 0; i < times; i++ {
    //     repeated += character
    // }
    return strings.Repeat(character, times)
}
