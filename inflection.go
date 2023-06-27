package humanizer

// Pluralize returns the plural form of the given word. If the word is already
// pluralized, the original word is returned. This function currently only
// supports a subset of English. If the word is not supported, the original
// word is returned. For example:
//
//     Pluralize("Person") // People
//     Pluralize("people") // people
func Pluralize(value string) string {
    return defaultVocabulary.pluralize(value)
}

// Singularize returns the singular form of the given word. If the word is
// already singular, the original word is returned. This function currently
// only supports a subset of English. If the word is not supported, the
// original word is returned. For example:
//
//     Singularize("People") // Person
//     Singularize("person") // person
func Singularize(value string) string {
    return defaultVocabulary.singularize(value)
}
