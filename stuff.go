package main

import (
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello fsdfs"))
}

func main() {
    http.HandleFunc("/", helloHandler)

    if err := http.ListenAndServe(":9000", nil); err != nil {
        log.Fatalln("ListenAndServer Error", err)
    }

}