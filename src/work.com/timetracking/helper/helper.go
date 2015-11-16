package helper

import (
    "fmt"
)

func PrintStringArrayForTables(name string, values []string) {
    fmt.Printf("Values for %s", name)
    if values != nil {
        for value := range values {
            fmt.Printf(": %s", values[value])
        }
        fmt.Printf(": LEN: %d\n", len(values))
    } else {
        fmt.Printf("no values")
    }
}

func PanicOnError(e error) {
    if e != nil {
        panic(e)
    }
}
