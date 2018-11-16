# Golang Metrics

## Description

This library contains helpers for easily instrumenting your code.

Currently it only supports [Prometheus](http://prometheus.io)

It exposes a simple wrapper around the the [promhttp.Handler](https://godoc.org/github.com/prometheus/client_golang/prometheus/promhttp#Handler)
as well as a selection of core metrics that follow the [RED Method](https://www.weave.works/blog/the-red-method-key-metrics-for-microservices-architecture/)

## Usage

### Core Metrics

To instrument an http handler with the core metrics you can do the following:

```go
package vertical

type Handler struct {}

var	metrics = golangmetrics.NewDefaultMetrics("valuecfg")

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) http.HandleFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        timer := prometheus.NewTimer(metrics.RequestDuration.WithLabelValues("index"))
	      metrics.RequestRate.WithLabelValues("index").Inc()
	      defer timer.ObserveDuration()

        someData, err := h.Worker.GoGetSomething()
	      if err != nil {
	          metrics.ErrorRate.WithLabelValues("index").Inc()
	          return errorHandler(w, r)
	      }

        w.Write(someData)
    }
}
```

## Metric Types

The library supports the 4 main metric types that enable you to expose data about your applications.

### Counter

A counter, very simply, is a single numeric value that can only increase. The wording that the prometheus docs uses is as follows:

> A counter is a cumulative metric that represents a single monotonically increasing counter whose value can only increase
> or be reset to zero on restart. For example, you can use a counter to represent the number of requests served, tasks completed, or errors.

### Gauge

A gauge is a single numeric value that can both go up and go down.

### Histogram

A histogram divides data into configurable buckets as well as an overall sum.

### Summary

A summary is very similar to a histogram in that it samples data and calculates configurable quantiles over time.
