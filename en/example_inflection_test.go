package en

import "fmt"

func ExamplePlural() {
    fmt.Println(Plural("Person"))
    // Output: People
}

func ExampleSingular() {
    fmt.Println(Singular("people"))
    // Output: person
}
