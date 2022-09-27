package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello, ชาวโลก")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Open your browser and go to http://localhost:8888")
    fmt.Println("Press Ctrl+C to exit...")
    log.Fatal(http.ListenAndServe(":8888", nil))
}
