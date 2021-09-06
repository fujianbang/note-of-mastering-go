package main

import (
	"fmt"
	"net/http"
)

var port = ":1443"

func Default(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is an example HTTPS server!\n")
}

func main() {
	http.HandleFunc("/", Default)
	fmt.Println("Listening to port number", port)

	if err := http.ListenAndServeTLS(port, "server.crt", "server.key", nil); err != nil {
		fmt.Println("ListenAndServeTLS: ", err)
		return
	}
}
