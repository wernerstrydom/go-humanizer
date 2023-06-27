package humanizer

import "testing"

func TestCamelCase(t *testing.T) {
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
                want: "deviceType",
            },
        )
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := CamelCase(tt.args.value); got != tt.want {
                    t.Errorf("CamelCase() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
