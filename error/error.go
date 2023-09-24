package error

import (
    "fmt"
    "os"
)

var HadError = false

func Error(line int, msg string) {
    Report(line, "", msg)
}

func Report(line int, where string, msg string) {
    fmt.Fprintf(os.Stderr, "[line %d] Error %s: %s\n", line, where, msg)
    HadError = true
}
