package main

import (
	//"errors"
	"fmt"
	//"io"
	"log"
	"net/http"
    //"os"
    "time"
)

func main() {
    fmt.Println("Booting Server")
    
    http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(writer, "<html><h1>Hello World</p></html>")
    })

    http.HandleFunc("/getcookie", func(writer http.ResponseWriter, r *http.Request) {
        expire := time.Now().Add(24*time.Hour)
        cookie := http.Cookie{
            Name: "user",
            Value: "password",
            Path: "/",
            Domain: "localhost",
            Expires: expire, 
            MaxAge: 24*60*60,
            Secure: false,
            HttpOnly: true,
        }
        http.SetCookie(writer, &cookie)
        fmt.Fprintf(writer, "<html><h1>Gave you a cookie!</p></html>")
    })

    http.HandleFunc("/removecookie", func(writer http.ResponseWriter, r *http.Request) {
        cookie := http.Cookie{
            Name: "user",
            Value: "",
            Path: "/",
            Expires: time.Now().Add(-100*time.Hour),
            MaxAge: -1,
            HttpOnly: true,
        }
        http.SetCookie(writer, &cookie)
        fmt.Fprintf(writer, "<html><h1>Cookie is gone!</p></html>")
    })

    log.Fatal(http.ListenAndServe(":8080",nil))
}

