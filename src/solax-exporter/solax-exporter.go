package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"solax-exporter/src/api"
)

type MetricsHandler struct {
	cloud_api *api.CloudAPI
}

func (m MetricsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.cloud_api != nil {
		cloud_response, err := m.cloud_api.Request()

		if err == nil {
			io.WriteString(w, "# HELP yield_total Total energy yield of inverter in KWh\n")
			io.WriteString(w, "# TYPE yield_total counter\n")
			io.WriteString(w, fmt.Sprintf("yield_total{sn=\"%s\"}=%f\n", m.cloud_api.SN, cloud_response.YieldTotal))
		}
	}
	io.WriteString(w, "#TODO metrics\n")
}

func root(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "#TODO root")
}

func main() {
	var port int
	var sn string
	var token_id string

	// Should be possible to provide more than one SN
	flag.StringVar(&sn, "sn", "", "Specify inverter SN.")
	flag.StringVar(&token_id, "token-id", "", "Specify Solax API token-id.")
	flag.IntVar(&port, "p", 9100, "Port to serve.")

	flag.Parse()

	metrics := MetricsHandler{
		api.MakeCloudApi(sn, token_id),
	}

	http.HandleFunc("/", root)
	http.Handle("/metrics", metrics)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Printf("Error %s\n", err)
	}
}
