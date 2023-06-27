package humanizer

import "strings"

// CamelCase converts a string to camel case, for example, "device type" to "deviceType"
// or "Device Type" to "deviceType", or "DeviceType" to "deviceType", or "deviceType" to "deviceType"
func CamelCase(value string) string {
    var result = PascalCase(value)
    if len(result) > 0 {
        return strings.ToLower(string(result[0])) + result[1:]
    }
    return result
}
