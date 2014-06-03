package main

import "log"
import "fmt"
import "github.com/linuturk/reddit"

func main() {
    items, err := reddit.Get("golang")
    if err != nil {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Println(item)
    }
}
