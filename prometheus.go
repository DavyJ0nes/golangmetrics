package golangmetrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// NewMetricsHandler is a simple wrapper around the prometheus provided handler
func NewMetricsHandler() http.Handler {
	return prometheus.Handler()
}

// DefaultMetrics contains the set of base metrics that should be added to each handler
type DefaultMetrics struct {
	RequestRate     *prometheus.GaugeVec
	RequestDuration *prometheus.HistogramVec
	ErrorRate       *prometheus.GaugeVec
}

// NewDefaultMetrics returns an initialised set of base metrics
func NewDefaultMetrics(handlerName string) DefaultMetrics {
	requestRate := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_request_rate", handlerName),
		Help: "The rate of requests per second",
	}, []string{"method"})

	requestDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: fmt.Sprintf("%s_handler_duration_seconds", handlerName),
		Help: "The request duration in seconds",
	}, []string{"method"})

	errorRate := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_request_error_rate", handlerName),
		Help: "The rate of errors per second",
	}, []string{"method"})

	prometheus.MustRegister(requestRate, requestDuration, errorRate)

	return DefaultMetrics{
		RequestRate:     requestRate,
		RequestDuration: requestDuration,
		ErrorRate:       errorRate,
	}
}
