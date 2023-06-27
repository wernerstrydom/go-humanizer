package humanizer

import (
    "strings"
    "unicode"
)

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
