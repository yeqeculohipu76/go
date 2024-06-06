package main

import (
    "fmt"
    "net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
    ch := make(chan string)
    for _, domain := range domains {
        go checkDomain(domain, ch)
    }

    for range domains {
        result := <-ch
        if result != "" {
            http.Redirect(w, r, result, http.StatusFound)
            return
        }
    }
    http.Error(w, "No domains are online.", http.StatusNotFound)
}

func main() {
    http.HandleFunc("/redirect", handler)
    fmt.Println("Server is running on port 3000")
    http.ListenAndServe(":3000", nil)
}
