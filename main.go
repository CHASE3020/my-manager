package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}



func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	 http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("public"))))


	log.Println("Server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
