package test

import (
	"bytes"
	"solax-exporter/src/metrics"
	"testing"
)

const expected_float_metric = `# HELP test_metric This is a test metric, in tests/s
# TYPE test_metric gauge
test_metric{sn="NOTAREALSN"}=10
`

func TestWriteFloatMetric(t *testing.T) {
	var w bytes.Buffer

	metrics.WriteFloatMetric(
		&w, metrics.Metric{
			Name: "test_metric",
			Help: "This is a test metric, in tests/s",
			Type: metrics.Gauge,
		}, "NOTAREALSN", 10.0)

	if w.String() != expected_float_metric {
		t.Errorf("Metric output error. Wanted %s, got %s", expected_float_metric, w.String())
	}
}
