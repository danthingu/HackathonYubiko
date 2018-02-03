package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<body><h1>Hello world!</h1></body>")
}

func loadPage(page string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filename := "./view/" + page
		body, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		w.Write(body)
	}
}

func main() {
	http.HandleFunc("/login", loadPage("login.html"))
	http.HandleFunc("/", handler)
	//http.HandleFunc("/intro/", loadPage("login.ht"))

	http.ListenAndServe(":8080", nil)
}
