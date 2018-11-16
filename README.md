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

import (
	"net/http"

	"github.com/davyj0nes/golangmetrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/someCompany/some"
)

type Handler struct {
	Metrics golangmetrics.CoreMetrics
	Worker 	some.Worker
}

func NewHandler() *Handler {
	return &Handler{
		Metrics: golangmetrics.NewDefaultMetrics("vertical")
		Worker:  some.InitWorker()
	}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) http.HandleFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        timer := prometheus.NewTimer(h.Metrics.RequestDuration.WithLabelValues("index"))
	    h.Metrics.RequestRate.WithLabelValues("index").Inc()
	    defer timer.ObserveDuration()

        someData, err := h.Worker.GoGetSomething()
	    if err != nil {
	        h.Metrics.ErrorRate.WithLabelValues("index").Inc()
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
