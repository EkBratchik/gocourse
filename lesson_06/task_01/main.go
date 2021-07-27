package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type handler struct{}

func (h handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	store := make(map[string]string)
	store["host:"] = (*req).Host
	store["user_agent:"] = (*req).UserAgent()
	store["request_uri:"] = (*req).RequestURI
	var header string
	for k, v := range (*req).Header {
		header += fmt.Sprintf("%s : %s\n", k, v)
	}
	store["headers:"] = header
	response, _ := json.Marshal(store)
	fmt.Fprintf(resp, string(response))
}
func main() {
	fmt.Println("Server is listening")
	var h handler
	http.ListenAndServe("localhost:8080", h)
}
