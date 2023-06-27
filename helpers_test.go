package humanizer

import (
    "fmt"
    "strings"
)

// join wraps each string in quotes and joins them with a comma and space.
func join(values []string) string {
    var sb strings.Builder
    for _, value := range values {
        sb.WriteString(fmt.Sprintf(`"%s", `, value))
    }

    return strings.Trim(strings.TrimSpace(sb.String()), ",")
}

var permutations2 = []string{
    "device type",
    "device_type",
    "device-type",
    "deviceType",
    "Device Type",
    "Device_Type",
    "Device-Type",
    "DeviceType",
    "$device Type",
    "$device Type$",
    "device$$type",
}
