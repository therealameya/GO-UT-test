package main

import (
    "bytes"
    "fmt"
    "testing"
)

func TestPrintAnything(t *testing.T) {
    tests := []struct {
        name  string
        value interface{}
    }{
        {"String", "Hello, World!"},
        {"Integer", 123},
        {"Float", 45.67},
        {"Boolean", true},
        {"Slice", []string{"apple", "banana", "cherry"}},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Capture the output
            var buf bytes.Buffer
            fmt.Fprint(&buf, tt.value)

            // Execute the function
            printAnything(tt.value)

            // Since printAnything does not return a value but prints to stdout,
            // this test is limited in its ability to verify the function's behavior.
            // Ideally, printAnything would return a value or write to an io.Writer passed as an argument.
        })
    }
}
