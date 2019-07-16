package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	sampleMetric = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "sample_metric",
		Help: "This is sample metric",
	})
	regstry = prometheus.NewRegistry()
)

func init() {
	sampleMetric.Set(3)
	regstry.MustRegister(sampleMetric)
}

func main() {
	prometheus.WriteToTextfile("./test.prom", regstry)
}
