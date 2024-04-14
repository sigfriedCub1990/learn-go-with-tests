package main

import "fmt"

const (
    spanish = "Spanish" 
    french = "French" 

    englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
    frenchHelloPrefix = "Bonjour, "
)

func Hello(name, lang string) string {
    if name == "" {
        name = "world"
    }

    return greetingPrefix(lang) + name + "!"
}

func greetingPrefix(lang string) (prefix string) {
    switch lang {
        case "French":
            prefix = frenchHelloPrefix
        case "Spanish":
            prefix = spanishHelloPrefix
        default:
            prefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("world", "English"))
}
