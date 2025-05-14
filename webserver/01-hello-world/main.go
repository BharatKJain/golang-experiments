package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Println("Running")
    })

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
        http.ServeFile(w,r,"index.html")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}