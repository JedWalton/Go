package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]
        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    r.Use(loggingMiddleware)

    log.Println("Server started on :8080")
    http.ListenAndServe(":8080", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
        next.ServeHTTP(w, r)
    })
}
