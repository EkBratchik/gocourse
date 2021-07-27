package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "main.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		name := r.FormValue("name")
		address := r.FormValue("address")

		cookies := &http.Cookie{
			Name:  "token",
			Value: name + ":" + address,
		}
		http.SetCookie(w, cookies)

		fmt.Fprintf(w, "Token = %s", cookies.Value)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	const port = 8080
	fmt.Printf("Launching server on port: %d \n\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
