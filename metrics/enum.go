package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func EnumValue(value bool) float64 {
	if value {
		return 1.0
	}
	return 0.0
}

type EnumCollector map[string]*prometheus.Desc

func (e EnumCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, desc := range e {
		ch <- desc
	}
}

func (e EnumCollector) Collect(ch chan<- prometheus.Metric, val string, labelValues ...string) {
	for v, desc := range e {
		ch <- prometheus.MustNewConstMetric(
			desc,
			prometheus.GaugeValue,
			EnumValue(v == val),
			labelValues...,
		)
	}
}

type Enum interface {
	Describe(chan<- *prometheus.Desc, string, string, prometheus.Labels)
	Collect(chan<- prometheus.Metric)
}
