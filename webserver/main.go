package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "parseForm() error:", err)
		return
	}
	fmt.Print("form post succesfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name %v", name)
	fmt.Fprintf(w, "Address %v", address)

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 NOt found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method NOt supported", http.StatusNotFound)
		return
	}
	fmt.Print("Hello")

}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", handleForm)
	http.HandleFunc("/hello", handleHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic("can't starrt server")
	}
	fmt.Println("Starting server at port 8080")
}
