package main

import "fmt"

// printAnything prints whatever is passed as a parameter
func printAnything(value interface{}) {
    fmt.Println(value)
}

func main() {
    // Example usage of printAnything
    printAnything("Hello, World!")
    printAnything(123)
    printAnything(45.67)
    printAnything(true)
    printAnything([]string{"apple", "banana", "cherry"})
}
