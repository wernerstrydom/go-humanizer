package en

func Plural(value string) string {
    return defaultVocabulary.pluralize(value)
}

func Singular(value string) string {
    return defaultVocabulary.singularize(value)
}