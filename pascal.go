package humanizer

import "strings"

// PascalCase converts a string to Pascal case. For example:
//
//     PascalCase("the brown fox jumps over the fence") // TheBrownFoxJumpsOverTheFence
//
func PascalCase(value string) string {
    words := ToWords(trim(value))
    if len(words) == 0 {
        return ""
    }
    var result strings.Builder
    for _, word := range words {
        if len(word) > 0 {
            result.WriteString(strings.ToUpper(string(word[0])))
            result.WriteString(word[1:])
        }
    }
    return result.String()
}
