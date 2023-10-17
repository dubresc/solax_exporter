package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func metrics(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "#TODO metrics")
}

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "#TODO root")
}

func main() {
	var port int
	var sn string

	flag.StringVar(&sn, "sn", "", "Specify inverter SN.")
	flag.IntVar(&port, "p", 9100, "Port to serve.")

	flag.Parse()

	http.HandleFunc("/", root)
	http.HandleFunc("/metrics", metrics)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
}
