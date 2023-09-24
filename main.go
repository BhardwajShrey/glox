package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
    "glox/error"
	"glox/scanner"
)

const LOGO = `

 ██▓     ▒█████  ▒██   ██▒
▓██▒    ▒██▒  ██▒▒▒ █ █ ▒░
▒██░    ▒██░  ██▒░░  █   ░
▒██░    ▒██   ██░ ░ █ █ ▒ 
░██████▒░ ████▓▒░▒██▒ ▒██▒
░ ▒░▓  ░░ ▒░▒░▒░ ▒▒ ░ ░▓ ░
░ ░ ▒  ░  ░ ▒ ▒░ ░░   ░▒ ░
  ░ ░   ░ ░ ░ ▒   ░    ░  
    ░  ░    ░ ░   ░    ░  
                          

`


const prompt = "> "

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
    data, err := os.ReadFile(path)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Filepath '%s' specified wrong. Please re-check...\n", path)
        os.Exit(1)
    }
    run(string(data))

    if error.HadError != false {
        os.Exit(65)
    }
}

func runPrompt() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Print(LOGO)
    fmt.Printf("Hello %s, welcome to Lox...\n\n", user.Username)
    
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Printf(prompt)
        isScanned := scanner.Scan()
        if !isScanned {
            // error while scanning tokens or EOF. stdin is however an always open file so no EOF unless Ctrl^C
            return
        }

        line := scanner.Text()  // returns most recent token scanned. Token here is a string of bytes (the line of code you just typed)
        if line == "" {
            break
        }
        run(line)

        error.HadError = false
    }
}

func run(source string) {
    s := scanner.New(source)
    tokens := s.ScanTokens()

    for _, t := range tokens {
        fmt.Println(t)
    }
}
