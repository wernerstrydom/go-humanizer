package humanizer

import "regexp"

func trim(value string) string {
    // use a regex to replace all leading non-alphanumeric characters with an empty string
    re := regexp.MustCompile(`^[^a-zA-Z0-9]+`)
    value = re.ReplaceAllString(value, "")

    // use a regex to replace all trailing non-alphanumeric characters with an empty string
    re = regexp.MustCompile(`[^a-zA-Z0-9]+$`)
    value = re.ReplaceAllString(value, "")

    // use a regex to replace all non-alphanumeric characters with a single space
    re = regexp.MustCompile(`[^a-zA-Z0-9]+`)
    value = re.ReplaceAllString(value, " ")

    // use a regex to replace all multiple spaces with a single space
    re = regexp.MustCompile(` +`)
    value = re.ReplaceAllString(value, " ")

    return value
}
