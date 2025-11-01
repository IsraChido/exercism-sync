package logs

import "fmt"
import "strings"
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range log {

        switch fmt.Sprintf("%U", v) {
        case "U+2757":
            return "recommendation"
        case "U+1F50D":
            return "search"
        case "U+2600":
            return "weather"
        }
    }

    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var logAsSlice []string
    
	for _, v := range log {
		if fmt.Sprintf("%U", v) == fmt.Sprintf("%U", oldRune) {
			logAsSlice = append(logAsSlice, string(newRune))
        } else {
            logAsSlice = append(logAsSlice, string(v))
        }
    }

    return strings.Join(logAsSlice, "")
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
    logLength := utf8.RuneCountInString(log)

    if logLength <= limit {
        return true
    } else {
        return false
    }
}
