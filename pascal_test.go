package humanizer

import (
    "log"
    "testing"
)

func TestPascalCase(t *testing.T) {
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
                want: "DeviceType",
            },
        )
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                got := PascalCase(tt.args.value)

                if got != tt.want {
                    t.Errorf("PascalCase(`%v`) returned `%v`, want `%v`", tt.args.value, got, tt.want)
                } else {
                    log.Printf("PascalCase(`%v`) returned `%v`, want `%v`\n", tt.args.value, got, tt.want)
                }
            },
        )
    }
}
