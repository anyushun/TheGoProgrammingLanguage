package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("command: ", os.Args[0])
    for i, arg := range os.Args[1:] {
        fmt.Println("    ", i, ": ", arg)
    }
}
