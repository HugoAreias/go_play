package main

import (
    "fmt"
    "net/http"
)

type String string

func (str String) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, str)
}

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}

func (s *Struct) ServeHTTP(
    w http.ResponseWriter,
    r *http.Request) {
    fmt.Fprint(w, s)
}

func (s *Struct) String() string {
    return fmt.Sprintf("%s%s%s", s.Greeting, s.Punct, s.Who)
}

func main() {
    var str String = "my string"
    http.ListenAndServe("localhost:4000", str)

    var s *Struct = &Struct{"Hello", ":", "Gophers!"}
    http.ListenAndServe("localhost:4000", s)

    http.Handle("/string", String("I'm a frayed knot."))
    http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
}