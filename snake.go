package humanizer

import "strings"

// SnakeCase converts a string to snake case. For example:
//
//    SnakeCase("The brown fox jumps over the fence") // the_brown_fox_jumps_over_the_fence
//
func SnakeCase(value string) string {
    words := ToWords(trim(value))
    if len(words) == 0 {
        return ""
    }
    var result strings.Builder
    for i, word := range words {
        if i > 0 {
            result.WriteString("_")
        }
        result.WriteString(strings.ToLower(word))
    }
    return result.String()
}
