package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/universe":
		fmt.Fprint(w, "Hello Universe ")
	default:
		fmt.Fprint(w, "<h1>Nowhere!</h1>")
	}
	fmt.Printf("What type of request comming %s\n", r.Method)
}

func plaintext(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Nowhere!</h1>")
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	//http.HandleFunc("/", plaintext)

	http.ListenAndServe("", nil)

}
