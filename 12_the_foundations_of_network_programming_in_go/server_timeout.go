package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler1(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	body := "The current time is"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", body)
	fmt.Fprintf(w, "<h2 align=\"center\"%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	port := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("Listening on http://0.0.0.0%s\n", port)
	} else {
		port = ":" + arguments[1]
		fmt.Printf("Listening on http://0.0.0.0%s\n", port)
	}

	m := http.NewServeMux()
	srv := &http.Server{
		Addr:         port,
		Handler:      m,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	m.HandleFunc("/time", timeHandler1)
	m.HandleFunc("/", myHandler1)

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
		return
	}
}
