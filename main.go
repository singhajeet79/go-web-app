package main

import (
    "net/http"
    "log"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/home.html")
    })
    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/about.html")
    })
    http.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/topics.html")
    })
    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/contact.html")
    })
    http.HandleFunc("/kubernetes-basics", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/kubernetes-basics.html")
    })
    http.HandleFunc("/advanced", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/advanced.html")
    })

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
