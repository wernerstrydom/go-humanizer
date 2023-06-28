package humanizer

import "testing"

func TestTitleCase(t *testing.T) {
    type args struct {
        value string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            name: "empty",
            args: args{
                value: "",
            },
            want: "",
        },
        {
            name: "one word",
            args: args{
                value: "word",
            },
            want: "Word",
        },
        {
            name: "two words",
            args: args{
                value: "two words",
            },
            want: "Two Words",
        },
        {
            name: "pascal case",
            args: args{
                value: "PascalCase",
            },
            want: "Pascal Case",
        },
        {
            name: "kebab case",
            args: args{
                value: "kebab-case",
            },
            want: "Kebab Case",
        },
        {
            name: "snake case",
            args: args{
                value: "snake_case",
            },
            want: "Snake Case",
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := TitleCase(tt.args.value); got != tt.want {
                    t.Errorf("TitleCase() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
