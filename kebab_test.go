package humanizer

import (
    "log"
    "testing"
)

func TestKebabCase(t *testing.T) {
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
    }

    for _, value := range permutations2 {
        tests = append(
            tests, struct {
                name string
                args args
                want string
            }{
                name: value,
                args: args{
                    value: value,
                },
                want: "device-type",
            },
        )
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := KebabCase(tt.args.value); got != tt.want {
                    t.Errorf("KebabCase(`%v`) = `%v`, want `%v`", tt.args.value, got, tt.want)
                } else {
                    log.Printf("KebabCase(`%v`) = `%v`, want `%v`\n", tt.args.value, got, tt.want)
                }
            },
        )
    }
}
