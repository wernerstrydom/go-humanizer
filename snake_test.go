package humanizer

import (
    "log"
    "testing"
)

func TestSnakeCase(t *testing.T) {
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
                want: "device_type",
            },
        )
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := SnakeCase(tt.args.value); got != tt.want {
                    t.Errorf("SnakeCase(`%v`) = `%v`, want `%v`", tt.args.value, got, tt.want)
                } else {
                    log.Printf("SnakeCase(`%v`) = `%v`, want `%v`\n", tt.args.value, got, tt.want)
                }
            },
        )
    }
}
