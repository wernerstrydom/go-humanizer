package humanizer

import "fmt"

func ExampleTitleCase() {
    fmt.Println(TitleCase("The brown fox jumps over the fence"))
    // Output: The Brown Fox Jumps Over The Fence
}

func ExampleTitleCase_pascalCase() {
    fmt.Println(TitleCase("TheBrownFoxJumpsOverTheFence"))
    // Output: The Brown Fox Jumps Over The Fence
}

func ExampleTitleCase_kebabCase() {
    fmt.Println(TitleCase("the-brown-fox-jumps-over-the-fence"))
    // Output: The Brown Fox Jumps Over The Fence
}
