package humanizer

import "strings"

// TitleCase converts a string to title case. For example:
//
//     TitleCase("the brown fox jumps over the fence") // The Brown Fox Jumps Over The Fence
//
func TitleCase(value string) string {
    words := ToWords(trim(value))
    if len(words) == 0 {
        return ""
    }

    var result strings.Builder
    for i, word := range words {
        if i > 0 {
            result.WriteString(" ")
        }
        result.WriteString(PascalCase(word))
    }
    return result.String()
}
