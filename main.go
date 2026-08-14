package main

import (
    "fmt"
    "os"
    "time"
)

func main() {
    fmt.Fprintln(os.Stderr, "sleeping...")
    time.Sleep(60 * time.Second)
    fmt.Fprintln(os.Stderr, "done")
}
