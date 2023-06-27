package humanizer

import (
    "log"
    "reflect"
    "testing"
)

func TestToWords(t *testing.T) {
    type args struct {
        value string
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        // tests for an empty string, one or two words, and strings seperated by spaces, underscores, dashes or other symbols, and strings with leading or trailing spaces, underscores, dashes or other symbols, and string that are already in pascal, camel, snake, kebab, or screaming case.
        {
            name: "empty",
            args: args{
                value: "",
            },
            want: []string{},
        },
        {
            name: "one word",
            args: args{
                value: "word",
            },
            want: []string{"word"},
        },
        {
            name: "two words",
            args: args{
                value: "two words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with spaces",
            args: args{
                value: "two words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with underscores",
            args: args{
                value: "two_words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with dashes",
            args: args{
                value: "two-words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with other symbols",
            args: args{
                value: "two*words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with leading spaces",
            args: args{
                value: "  two words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with trailing spaces",
            args: args{
                value: "two words  ",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with leading underscores",
            args: args{
                value: "__two words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with trailing underscores",
            args: args{
                value: "two words__",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with leading dashes",
            args: args{
                value: "--two words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with trailing dashes",
            args: args{
                value: "two words--",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with leading other symbols",
            args: args{
                value: "**two words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words with trailing other symbols",
            args: args{
                value: "two words**",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words in pascal case",
            args: args{
                value: "TwoWords",
            },
            want: []string{"Two", "Words"},
        },
        {
            name: "two words in camel case",
            args: args{
                value: "twoWords",
            },
            want: []string{"two", "Words"},
        },
        {
            name: "two words in snake case",
            args: args{
                value: "two_words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words in kebab case",
            args: args{
                value: "two-words",
            },
            want: []string{"two", "words"},
        },
        {
            name: "two words in screaming case",
            args: args{
                value: "TWO_WORDS",
            },
            want: []string{"TWO", "WORDS"},
        },
        {
            name: "a sentence",
            args: args{
                value: "The brown fox jumped over the lazy dog.",
            },
            want: []string{"The", "brown", "fox", "jumped", "over", "the", "lazy", "dog"},
        },
        {
            name: "a sentence with symbols",
            args: args{
                value: "The brown fox - which I love very much - jumped over the lazy dog.",
            },
            want: []string{
                "The",
                "brown",
                "fox",
                "which",
                "I",
                "love",
                "very",
                "much",
                "jumped",
                "over",
                "the",
                "lazy",
                "dog",
            },
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := ToWords(tt.args.value); !reflect.DeepEqual(got, tt.want) {
                    t.Errorf(
                        "ToWords(`%v`) returned [%v], and we expected [%v]",
                        tt.args.value,
                        join(got),
                        join(tt.want),
                    )
                } else {
                    log.Printf("ToWords(`%v`)returned [%v]", tt.args.value, join(got))
                }
            },
        )
    }
}
