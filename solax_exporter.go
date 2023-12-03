package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"solax-exporter/metrics"
	"solax-exporter/solax"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type CloudCollector struct {
	cloud_api *solax.CloudApiRequester
}

var (
	acPowerDesc        *prometheus.Desc       = prometheus.NewDesc("ac_power", "Current load on the inverter in W", nil, nil)
	yieldTodayDesc     *prometheus.Desc       = prometheus.NewDesc("yield_today", "Total energy yield of inverter for the day in kWh", nil, nil)
	yieldTotalDesc     *prometheus.Desc       = prometheus.NewDesc("yield_total", "Total energy yield of inverter in kWh", nil, nil)
	feedInPowerDesc    *prometheus.Desc       = prometheus.NewDesc("feed_in_power", "Current load on the grid in W", nil, nil)
	feedInEnergyDesc   *prometheus.Desc       = prometheus.NewDesc("feed_in_energy", "Total energy fed to grid in kWh", nil, nil)
	comsumeEnergyDesc  *prometheus.Desc       = prometheus.NewDesc("consume_energy", "Total energy consumed from grid in kWh", nil, nil)
	feedInPowerM2      *prometheus.Desc       = prometheus.NewDesc("feed_in_power_m2", "Current load on the grid in W", nil, nil)
	socDesc            *prometheus.Desc       = prometheus.NewDesc("soc", "Current state of charge of the battery in %", nil, nil)
	peps1Desc          *prometheus.Desc       = prometheus.NewDesc("peps1", "", nil, nil)
	peps2Desc          *prometheus.Desc       = prometheus.NewDesc("peps2", "", nil, nil)
	peps3Desc          *prometheus.Desc       = prometheus.NewDesc("peps3", "", nil, nil)
	batPowerDesc       *prometheus.Desc       = prometheus.NewDesc("bat_power", "Current load on the battery in W", nil, nil)
	inverterTypeDesc   *metrics.EnumCollector = solax.DescribeInverterTypeCode("inverter_type", nil, nil)
	inverterStatusDesc *metrics.EnumCollector = solax.DescribeInverterStatusCode("inverter_status", nil, nil)
	powerDC1Desc       *prometheus.Desc       = prometheus.NewDesc("power_dc1", "", nil, nil)
	powerDC2Desc       *prometheus.Desc       = prometheus.NewDesc("power_dc2", "", nil, nil)
	powerDC3Desc       *prometheus.Desc       = prometheus.NewDesc("power_dc3", "", nil, nil)
	powerDC4Desc       *prometheus.Desc       = prometheus.NewDesc("power_dc4", "", nil, nil)
)

func (g *CloudCollector) Describe(ch chan<- *prometheus.Desc) {
	if g.cloud_api != nil {
		ch <- acPowerDesc
		ch <- yieldTodayDesc
		ch <- yieldTotalDesc
		ch <- feedInPowerDesc
		ch <- feedInEnergyDesc
		ch <- comsumeEnergyDesc
		ch <- feedInPowerM2
		ch <- socDesc
		ch <- peps1Desc
		ch <- peps2Desc
		ch <- peps3Desc
		inverterTypeDesc.Describe(ch)
		inverterStatusDesc.Describe(ch)
		ch <- batPowerDesc
		ch <- powerDC1Desc
		ch <- powerDC2Desc
		ch <- powerDC3Desc
		ch <- powerDC4Desc
	}
}

func (g *CloudCollector) Collect(ch chan<- prometheus.Metric) {
	if g.cloud_api != nil {
		cloud_response, err := g.cloud_api.Request()

		if err != nil {
			fmt.Printf("Reuqest error %s\n", err)
			return
		}

		ch <- prometheus.MustNewConstMetric(acPowerDesc, prometheus.GaugeValue, cloud_response.ACPower)
		ch <- prometheus.MustNewConstMetric(yieldTodayDesc, prometheus.CounterValue, cloud_response.YieldToday)
		ch <- prometheus.MustNewConstMetric(yieldTotalDesc, prometheus.CounterValue, cloud_response.YieldTotal)
		ch <- prometheus.MustNewConstMetric(feedInPowerDesc, prometheus.GaugeValue, cloud_response.FeedInPower)
		ch <- prometheus.MustNewConstMetric(feedInEnergyDesc, prometheus.CounterValue, cloud_response.FeedInEnergy)
		ch <- prometheus.MustNewConstMetric(comsumeEnergyDesc, prometheus.CounterValue, cloud_response.ConsumeEnergy)
		ch <- prometheus.MustNewConstMetric(feedInPowerM2, prometheus.GaugeValue, cloud_response.FeedInPowerM2)
		ch <- prometheus.MustNewConstMetric(socDesc, prometheus.GaugeValue, cloud_response.SOC)
		if cloud_response.Peps1 != nil {
			ch <- prometheus.MustNewConstMetric(peps1Desc, prometheus.GaugeValue, *cloud_response.Peps1)
		}
		if cloud_response.Peps2 != nil {
			ch <- prometheus.MustNewConstMetric(peps2Desc, prometheus.GaugeValue, *cloud_response.Peps2)
		}
		if cloud_response.Peps2 != nil {
			ch <- prometheus.MustNewConstMetric(peps3Desc, prometheus.GaugeValue, *cloud_response.Peps3)
		}
		inverterTypeDesc.Collect(ch, cloud_response.InverterType.String())
		inverterStatusDesc.Collect(ch, cloud_response.InverterStatus.String())
		ch <- prometheus.MustNewConstMetric(batPowerDesc, prometheus.GaugeValue, cloud_response.BatPower)
		if cloud_response.PowerDC1 != nil {
			ch <- prometheus.MustNewConstMetric(powerDC1Desc, prometheus.GaugeValue, *cloud_response.PowerDC1)
		}
		if cloud_response.PowerDC2 != nil {
			ch <- prometheus.MustNewConstMetric(powerDC2Desc, prometheus.GaugeValue, *cloud_response.PowerDC2)
		}
		if cloud_response.PowerDC3 != nil {
			ch <- prometheus.MustNewConstMetric(powerDC3Desc, prometheus.GaugeValue, *cloud_response.PowerDC3)
		}
		if cloud_response.PowerDC4 != nil {
			ch <- prometheus.MustNewConstMetric(powerDC4Desc, prometheus.GaugeValue, *cloud_response.PowerDC4)
		}

	}
}

func main() {
	var port int
	var sn string
	var token_id string

	flag.StringVar(&sn, "sn", "", "Specify inverter SN.")
	flag.StringVar(&token_id, "token-id", "", "Specify Solax API token-id.")
	flag.IntVar(&port, "p", 9100, "Port to serve.")

	flag.Parse()

	r := prometheus.NewRegistry()
	r.MustRegister(&CloudCollector{
		solax.MakeCloudApiRequester(sn, token_id),
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Solax Exporter</title></head>
             <body>
             <h1>Solax Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`),
		)
	})
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}
}
