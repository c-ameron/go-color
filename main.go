package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    keys := r.URL.Query()
    hexColor := keys.Get("hex")
    fmt.Fprintf(w, "query %s", hexColor)
}

func main() {
    http.HandleFunc("/convert", handler)
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}