package humanizer

import "fmt"

func ExamplePlural() {
    fmt.Println(Pluralize("Person"))
    // Output: People
}

func ExampleSingular() {
    fmt.Println(Singularize("people"))
    // Output: person
}
