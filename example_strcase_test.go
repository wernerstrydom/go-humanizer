package humanizer

import "fmt"

func ExamplePascalCase() {
    fmt.Println(PascalCase("hello world"))
    // Output: HelloWorld
}

func ExampleCamelCase() {
    fmt.Println(CamelCase("hello world"))
    // Output: helloWorld
}

func ExampleSnakeCase() {
    fmt.Println(SnakeCase("hello world"))
    // Output: hello_world
}

func ExampleKebabCase() {
    fmt.Println(KebabCase("hello world"))
    // Output: hello-world
}
