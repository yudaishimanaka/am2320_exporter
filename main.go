package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/oltoko/go-am2320"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "am2320"
)

type Collector struct{}

var (
	temperatureGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "temperature_gauge",
			Help:      "temperature gauge help",
		})

	humiditiyGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "humiditiy_gauge",
			Help:      "humiditiy gauge help",
		})
)

func (c Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- temperatureGauge.Desc()
	ch <- humiditiyGauge.Desc()
}

func (c Collector) Collect(ch chan<- prometheus.Metric) {
	// Read and get tmp and hum for AM2320
	am2320 := am2320.Create(am2320.DefaultI2CAddr)
	res, err := am2320.Read()
	if err != nil {
		log.Fatalln("Failed to read from AM2320", err)
	}

	ch <- prometheus.MustNewConstMetric(
		temperatureGauge.Desc(),
		prometheus.GaugeValue,
		float64(res.Temperature),
	)

	ch <- prometheus.MustNewConstMetric(
		humiditiyGauge.Desc(),
		prometheus.GaugeValue,
		float64(res.Humidity),
	)
}

func main() {
	var listenAddress = flag.String(
		"listen-address",
		":9700",
		"The address to listen on for HTTP requests.",
	)
	flag.Parse()

	var c Collector
	prometheus.MustRegister(c)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}
