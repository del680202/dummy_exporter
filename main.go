package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

type Exporter struct {
	gauge    prometheus.Gauge
	gaugeVec prometheus.GaugeVec
}

func NewExporter(metricsPrefix string) *Exporter {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "gauge_metric",
		Help:      "This is a dummy gauge metric"})

	gaugeVec := *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsPrefix,
		Name:      "gauge_vec_metric",
		Help:      "This is a dummy gauga vece metric"},
		[]string{"myLabel"})

	return &Exporter{
		gauge:    gauge,
		gaugeVec: gaugeVec,
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.gauge.Set(float64(0))
	e.gaugeVec.WithLabelValues("hello").Set(float64(0))
	e.gauge.Collect(ch)
	e.gaugeVec.Collect(ch)
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.gauge.Describe(ch)
	e.gaugeVec.Describe(ch)
}

func main() {
	fmt.Println(`
  This is a dummy example of prometheus exporter
  Access: http://127.0.0.1:8081
  `)

	// Define parameters
	metricsPath := "/metrics"
	listenAddress := ":8081"
	metricsPrefix := "dummy"

	// Register dummy exporter, not necessary
	exporter := NewExporter(metricsPrefix)
	prometheus.MustRegister(exporter)

	// Launch http service
	http.Handle(metricsPath, prometheus.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Dummy Exporter</title></head>
             <body>
             <h1>Dummy Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	fmt.Println(http.ListenAndServe(listenAddress, nil))
}
