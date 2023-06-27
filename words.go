package humanizer

import (
    "strings"
    "unicode"
)

// ToWords converts a string to a slice of words. For example:
//
//    ToWords("The brown fox jumps over the fence") // []string{"The", "brown", "fox", "jumps", "over", "the", "fence"}
//    ToWords("the_brown_fox_jumps_over_the_fence") // []string{"the", "brown", "fox", "jumps", "over", "the", "fence"}
//    ToWords("the-brown-fox-jumps-over-the-fence") // []string{"the", "brown", "fox", "jumps", "over", "the", "fence"}
//    ToWords("theBrownFoxJumpsOverTheFence") // []string{"the", "Brown", "Fox", "Jumps", "Over", "The", "Fence"}
//    ToWords("TheBrownFoxJumpsOverTheFence") // []string{"The", "Brown", "Fox", "Jumps", "Over", "The", "Fence"}
func ToWords(value string) []string {
    if value == "" {
        return []string{}
    }
    var result []string
    upper := false
    sb := strings.Builder{}
    for _, c := range value {
        if (unicode.IsLetter(c) && unicode.IsUpper(c) && !upper) || unicode.IsDigit(c) || unicode.IsSymbol(c) || unicode.IsPunct(c) || unicode.IsSpace(c) {
            if sb.Len() > 0 {
                result = append(result, sb.String())
                sb.Reset()
            }
            upper = true
        }

        if !unicode.IsLetter(c) || !unicode.IsUpper(c) {
            upper = false
        }

        if !unicode.IsSymbol(c) && !unicode.IsPunct(c) && !unicode.IsSpace(c) {
            sb.WriteRune(c)
        }
    }
    if sb.Len() > 0 {
        result = append(result, sb.String())
        sb.Reset()
    }
    return result
}
