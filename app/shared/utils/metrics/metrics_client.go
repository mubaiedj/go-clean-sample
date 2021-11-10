package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	testMetrics *prometheus.CounterVec
	errors      *prometheus.CounterVec
}

const APP = "goCleanSample"

var metricsCli *metrics

func init() {
	var testMetrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_test_metrics", APP),
		Help: "Total of times called this test metric",
	}, []string{"event"})

	var errors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: fmt.Sprintf("%s_custom_errors", APP),
		Help: "Total of custom errors per kind",
	}, []string{"kind"})
	prometheus.MustRegister(testMetrics, errors)

	metricsCli = &metrics{
		testMetrics: testMetrics,
		errors:      errors,
	}
}

func IncrementTestMetrics(event string) { metricsCli.IncrementTestMetrics(event) }
func (metrics *metrics) IncrementTestMetrics(event string) {
	metrics.testMetrics.With(prometheus.Labels{"event": event}).Inc()
}

func IncrementErrors(kind string) { metricsCli.IncrementErrors(kind) }
func (metrics *metrics) IncrementErrors(kind string) {
	metrics.errors.With(prometheus.Labels{"kind": kind}).Inc()
}
