package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"solax-exporter/src/api"
)

type MetricsHandler struct {
	cloud_api api.CloudAPI
}

func (m MetricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "#TODO metrics")
}

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "#TODO root")
}

func main() {
	var port int
	var sn string
	var token_id string

	flag.StringVar(&sn, "sn", "", "Specify inverter SN.")
	flag.StringVar(&token_id, "token-id", "", "Specify Solax API token-id.")
	flag.IntVar(&port, "p", 9100, "Port to serve.")

	flag.Parse()

	metrics := MetricsHandler{
		api.CloudAPI{
			SN:      &sn,
			TokenID: &token_id,
		},
	}

	http.HandleFunc("/", root)
	http.Handle("/metrics", metrics)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
}
