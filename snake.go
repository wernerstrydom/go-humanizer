package humanizer

import "strings"

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
