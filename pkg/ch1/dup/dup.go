package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string]int)
    files := make(map[string]string)
    for _, fm := range os.Args[1:] {
        f, err := os.Open(fm)
        if err != nil {
            fmt.Printf("read file: %s with error: %s\n", fm, err.Error())
            continue
        }
        defer f.Close()
        input := bufio.NewScanner(f)
        for input.Scan() {
            counts[input.Text()]++
            if counts[input.Text()] == 2 {
                files[input.Text()] = fm
            }
        }
    }
    for v, fm := range files {
        fmt.Printf("%20s %3d %s\n", fm, counts[v], v)
    }
}
