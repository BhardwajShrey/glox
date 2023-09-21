package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]

    if len(args) > 1 {
        fmt.Printf("Usage: glox [script]\n")
        os.Exit(64)
    } else if len(args) == 1 {
        runFile(args[0])
    } else {
        runPrompt()
    }
}

func runFile(path string) {
    fmt.Printf("Running on %s...\n", path)
    // actually do something
}

func runPrompt() {
    // this is the REPL
    fmt.Printf("Welcome to Lox. Start typing away...\n")
    fmt.Printf("> ")
    fmt.Printf("\n")
    // do something here as well
}
