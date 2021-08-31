package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	port := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number:", port)
	} else {
		port = ":" + arguments[1]
	}

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println(err)
	}
}
