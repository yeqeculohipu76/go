package main

import (
    "fmt"
    "net/http"
    "time"
)

var domains = []string{
    "http://domain1.com",
    "http://domain2.com",
    "http://domain3.com",
}

func checkDomain(domain string, ch chan string) {
    _, err := http.Get(domain)
    if err == nil {
        ch <- domain
    } else {
        ch <- ""
    }
}

func main() {
    ch := make(chan string)
    for _, domain := range domains {
        go checkDomain(domain, ch)
    }

    for range domains {
        result := <-ch
        if result != "" {
            fmt.Printf("Redirecting to %s\n", result)
            return
        }
    }
    fmt.Println("No domains are online.")
}
