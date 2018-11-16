package golangmetrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

// CoreMetrics contains the set of base metrics that should be added to each handler
type CoreMetrics struct {
	RequestRate     *prometheus.GaugeVec
	RequestDuration *prometheus.HistogramVec
	ErrorRate       *prometheus.GaugeVec
}

// NewCoreMetrics returns an initialised set of base metrics
func NewCoreMetrics(appName, handlerName string) CoreMetrics {
	requestRate := NewGauge(
		appName,
		fmt.Sprintf("%s_request_rate", handlerName),
		"The rate of requests per second",
		[]string{"method"},
	)

	requestDuration := NewHistogram(
		appName,
		fmt.Sprintf("%s_handler_duration_seconds", handlerName),
		"The request duration in seconds",
		[]string{"method"},
	)

	errorRate := NewGauge(
		appName,
		fmt.Sprintf("%s_request_error_rate", handlerName),
		"The rate of errors per second",
		[]string{"method"},
	)

	prometheus.MustRegister(requestRate, requestDuration, errorRate)

	return CoreMetrics{
		RequestRate:     requestRate,
		RequestDuration: requestDuration,
		ErrorRate:       errorRate,
	}
}

// NewMetricsHandler is a simple wrapper around the prometheus provided handler
func NewMetricsHandler() http.Handler {
	return prometheus.Handler()
}

// NewCounter wraps prometheus.NewCounterVec
func NewCounter(namespace, name, help string, labels []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labels)
}

// NewGauge wraps prometheus.NewGaugeVec
func NewGauge(namespace, name, help string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labels)
}

// NewHistogram wraps prometheus.NewHistogramVec
func NewHistogram(namespace, name, help string, labels []string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labels)
}

// NewSummary wraps prometheus.NewSummaryVec
func NewSummary(namespace, name, help string, labels []string) *prometheus.SummaryVec {
	return prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	}, labels)
}
