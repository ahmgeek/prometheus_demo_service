package main

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	lastSuccess = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "batch",
			Name:      "last_success_timestamp_seconds",
			Help:      "The Unix timestamp in seconds since the last successful demo batch job completion.",
		},
	)
	lastRun = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "batch",
			Name:      "last_run_timestamp_seconds",
			Help:      "The Unix timestamp in seconds since the last demo batch job run.",
		},
	)
	lastDuration = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "batch",
			Name:      "last_run_duration_seconds",
			Help:      "The duration in seconds of the last batch job run.",
		},
	)
	processedRecords = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "batch",
			Name:      "last_run_processed_bytes",
			Help:      "The number of records processed by the demo batch job in the last run.",
		})
)

func init() {
	prometheus.MustRegister(lastSuccess)
	prometheus.MustRegister(lastRun)
	prometheus.MustRegister(lastDuration)
	prometheus.MustRegister(processedRecords)
}

func runBatchJobs(interval time.Duration, duration time.Duration, failureRatio float64) {
	lastTime := float64(time.Now().UnixNano()) / 1e9
	lastRecords := 1000
	ticker := time.NewTicker(interval)

	for {
		time.Sleep(duration + time.Second - time.Duration((rand.Int()%2000))*time.Microsecond)

		now := float64(time.Now().UnixNano()) / 1e9
		if rand.Float64() > failureRatio {
			lastSuccess.Set(now)
			lastRecords += 25 - rand.Int()%50
			processedRecords.Set(float64(lastRecords))
		}

		lastRun.Set(now)

		lastDuration.Set(float64(now - lastTime))
		lastTime = now

		<-ticker.C
	}
}
