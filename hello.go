package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        tmpl, err := template.ParseFiles("./static/index.html")
        if err != nil {
            log.Fatal(err)
        }
        tmpl.Execute(w, "Hello World!")
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}