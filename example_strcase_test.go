package humanizer

import "fmt"

func ExamplePascalCase() {
    fmt.Println(PascalCase("The brown fox jumps over the fence"))
    // Output: TheBrownFoxJumpsOverTheFence
}

func ExamplePascalCase_from_camelCase() {
    fmt.Println(PascalCase("theBrownFoxJumpsOverTheFence"))
    // Output: TheBrownFoxJumpsOverTheFence
}

func ExamplePascalCase_from_kebabCase() {
    fmt.Println(PascalCase("the-brown-fox-jumps-over-the-fence"))
    // Output: TheBrownFoxJumpsOverTheFence
}

func ExamplePascalCase_from_snakeCase() {
    fmt.Println(PascalCase("the_brown_fox_jumps_over_the_fence"))
    // Output: TheBrownFoxJumpsOverTheFence
}

func ExamplePascalCase_trimsSpace() {
    fmt.Println(PascalCase("   The brown fox jumps    over the fence"))
    // Output: TheBrownFoxJumpsOverTheFence
}

func ExampleCamelCase() {
    fmt.Println(CamelCase("The brown fox jumps over the fence"))
    // Output: theBrownFoxJumpsOverTheFence
}

func ExampleCamelCase_from_pascalCase() {
    fmt.Println(CamelCase("TheBrownFoxJumpsOverTheFence"))
    // Output: theBrownFoxJumpsOverTheFence
}

func ExampleCamelCase_from_kebabCase() {
    fmt.Println(CamelCase("the-brown-fox-jumps-over-the-fence"))
    // Output: theBrownFoxJumpsOverTheFence
}

func ExampleCamelCase_from_snakeCase() {
    fmt.Println(CamelCase("the_brown_fox_jumps_over_the_fence"))
    // Output: theBrownFoxJumpsOverTheFence
}

func ExampleCamelCase_trimsSpace() {
    fmt.Println(CamelCase("   The brown fox jumps    over the fence"))
    // Output: theBrownFoxJumpsOverTheFence
}

func ExampleSnakeCase() {
    fmt.Println(SnakeCase("The brown fox jumps over the fence"))
    // Output: the_brown_fox_jumps_over_the_fence
}

func ExampleSnakeCase_from_pascalCase() {
    fmt.Println(SnakeCase("TheBrownFoxJumpsOverTheFence"))
    // Output: the_brown_fox_jumps_over_the_fence
}

func ExampleSnakeCase_from_camelCase() {
    fmt.Println(SnakeCase("theBrownFoxJumpsOverTheFence"))
    // Output: the_brown_fox_jumps_over_the_fence
}

func ExampleSnakeCase_from_kebabCase() {
    fmt.Println(SnakeCase("the-brown-fox-jumps-over-the-fence"))
    // Output: the_brown_fox_jumps_over_the_fence
}

func ExampleSnakeCase_trimsSpace() {
    fmt.Println(SnakeCase("   The brown fox jumps    over the fence"))
    // Output: the_brown_fox_jumps_over_the_fence
}

func ExampleKebabCase() {
    fmt.Println(KebabCase("The brown fox jumps over the fence"))
    // Output: the-brown-fox-jumps-over-the-fence
}

func ExampleKebabCase_from_pascalCase() {
    fmt.Println(KebabCase("TheBrownFoxJumpsOverTheFence"))
    // Output: the-brown-fox-jumps-over-the-fence
}

func ExampleKebabCase_from_camelCase() {
    fmt.Println(KebabCase("theBrownFoxJumpsOverTheFence"))
    // Output: the-brown-fox-jumps-over-the-fence
}

func ExampleKebabCase_from_snakeCase() {
    fmt.Println(KebabCase("the_brown_fox_jumps_over_the_fence"))
    // Output: the-brown-fox-jumps-over-the-fence
}

func ExampleKebabCase_trimsSpace() {
    fmt.Println(KebabCase("   The brown fox jumps    over the fence"))
    // Output: the-brown-fox-jumps-over-the-fence
}
