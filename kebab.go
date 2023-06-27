package humanizer

import "strings"

// KebabCase converts a string to kebab case. For example:
//
//     KebabCase("The brown fox jumps over the fence") // the-brown-fox-jumps-over-the-fence
//
func KebabCase(value string) string {
    words := ToWords(trim(value))
    if len(words) == 0 {
        return ""
    }
    var result strings.Builder
    for i, word := range words {
        if i > 0 {
            result.WriteString("-")
        }
        result.WriteString(strings.ToLower(word))
    }
    return result.String()
}
