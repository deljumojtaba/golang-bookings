package main

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Home Page"))
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is About Page"))
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(":8080", nil)
}
