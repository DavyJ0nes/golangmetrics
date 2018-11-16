package golangmetrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// NewMetricsHandler is a simple wrapper around the prometheus provided handler
func NewMetricsHandler() http.Handler {
	return prometheus.Handler()
}

// NewCounter is a simple wrapper that creates a Prometheus Counter
func NewCounter(name, help string, labels []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		}, labels,
	)
}

// NewGauge is a simple wrapper that creates a Prometheus Gauge
func NewGauge(name, help string) {}

// NewHistogram is a simple wrapper that creates a Prometheus Histogram
func NewHistogram(name, help string) {}

// NewSummary is a simple wrapper that creates a Prometheus Summary
func NewSummary(name, help string) {}

// DefaultMetrics contains the set of base metrics that should be added to each handler
type DefaultMetrics struct {
	RequestRate     *prometheus.GaugeVec
	RequestDuration *prometheus.HistogramVec
	ErrorRate       *prometheus.GaugeVec
}

// NewDefaultMetrics returns an initialised set of base metrics
func NewDefaultMetrics(handlerName string) *DefaultMetrics {
	requestRate := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_request_rate", handlerName),
		Help: "The rate of requests per second",
	}, []string{"method"})

	requestDuration := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: fmt.Sprintf("%s_handler_duration_seconds", handlerName),
		Help: "The request duration in seconds",
	}, []string{"method"})

	errorRate := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: fmt.Sprintf("%s_request_error_rate", handlerName),
		Help: "The rate of errors per second",
	}, []string{"method"})

	return &DefaultMetrics{
		RequestRate:     requestRate,
		RequestDuration: requestDuration,
		ErrorRate:       errorRate,
	}
}
