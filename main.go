package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/form", handleForm)

	fmt.Println("server has started on port : 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleForm(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing the form %v", err)
		return
	}
	fmt.Printf("form post request success")
	name := r.FormValue("name")
	fmt.Fprintf(w, "%s \n", name)

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "METHOD NOT ALLOWED", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello")
}
