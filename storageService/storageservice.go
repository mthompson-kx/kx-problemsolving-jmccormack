package main

import (
    "fmt"
    "net/http"
    "flag"
)

const port = "8000"

func main() {

    debugPtr := flag.Bool("debug", false, "debug info")
    flag.Parse()
    debug := *debugPtr

    if debug {
        fmt.Println("**********************************")
        fmt.Println("Starting File Server")
        fmt.Println("File Server up")
        fmt.Printf("listening on http://localhost:%s/data.json \n", port)
        fmt.Printf("healthcheck on http://localhost:%s/health \n", port)
        fmt.Println("**********************************")
    }

    http.Handle("/data.json", http.FileServer(http.Dir("./static")))

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "ok")
    })

    http.ListenAndServe(fmt.Sprint(":", port), nil)
}
