package metrics

import (
	"fmt"
	"io"
)

type MetricType string

const (
	Counter MetricType = "counter"
	Gauge   MetricType = "gauge"
)

type Metric struct {
	Name string
	Help string
	Type MetricType
}

var (
	YieldTotal Metric = Metric{
		Name: "yield_total",
		Help: "Total energy yield of inverter in kWh",
		Type: Counter,
	}
)

func WriteFloatMetric(w io.Writer, m Metric, sn string, v float64) {
	io.WriteString(w, fmt.Sprintf("# HELP %s %s\n", m.Name, m.Help))
	io.WriteString(w, fmt.Sprintf("# TYPE %s %s\n", m.Name, m.Type))
	io.WriteString(w, fmt.Sprintf("%s{sn=\"%s\"}=%g", m.Name, sn, v))
}
