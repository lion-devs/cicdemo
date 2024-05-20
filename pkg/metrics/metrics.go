package metrics

import "github.com/penglongli/gin-metrics/ginmetrics"

func GetMonitor() *ginmetrics.Monitor {
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")

	// +optional set slow time, default 5s
	m.SetSlowTime(10)

	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	gaugeMetric := &ginmetrics.Metric{
		Type:        ginmetrics.Gauge,
		Name:        "example_gauge_metric",
		Description: "an example of gauge type metric",
		Labels:      []string{"label1"},
	}

	// Add metric to global monitor object
	_ = ginmetrics.GetMonitor().AddMetric(gaugeMetric)
	_ = ginmetrics.GetMonitor().GetMetric("example_gauge_metric").SetGaugeValue([]string{"label-test"}, 8.7)

	return m
}
