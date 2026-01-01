package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"go.szostok.io/version"
	"go.szostok.io/version/printer"
)

//nolint:funlen
func main() {

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:05",
	})

	var scrapeTimeout time.Duration
	var addr string
	var measurementsURL string
	var v bool

	flag.DurationVar(&scrapeTimeout, "timeout", 30*time.Second, "timeout duration between measurements")
	flag.StringVar(&addr, "http-addr", ":9500", "address for the Prometheus server")
	flag.StringVar(&measurementsURL, "measurements-url", "http://airgradient.local/measures/current", "URL to fetch AirGradient measurements from")
	flag.BoolVar(&v, "version", false, "print version")
	flag.BoolVar(&v, "v", false, "print version")
	flag.Parse()

	if addr == "" || measurementsURL == "" || scrapeTimeout == 0 {
		log.Fatal("addr, url, and timeout are required")
	}

	if v {
		var verPrinter = printer.New()
		var info = version.Get()
		if err := verPrinter.PrintInfo(os.Stdout, info); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	go func() {
		log.Infof("Starting Prometheus server on %s", addr)
		http.Handle("/metrics", promhttp.Handler())
		srv := &http.Server{
			Addr:         addr,
			Handler:      nil,
			ReadTimeout:  time.Second * 2,
			WriteTimeout: time.Second * 2,
		}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		<-sigChan
		os.Exit(0)
	}()

	for {
		measurements, err := getMeasurements(measurementsURL)
		if err != nil {
			log.Warn(err)
			continue
		}

		updateMetrics(measurements)

		time.Sleep(scrapeTimeout)
	}
}

// getMeasurements fetches the current measurements from the AirGradient device.
//
// we only use named returns here for the deferred error handling
//
//nolint:nonamedreturns
func getMeasurements(measurementsURL string) (measurements AirGradientMeasures, err error) {
	//nolint:gosec
	resp, err := http.Get(measurementsURL)
	if err != nil {
		return measurements, fmt.Errorf("failed to fetch measurements: %w", err)
	}
	defer func() {
		err = resp.Body.Close()
	}()

	err = json.NewDecoder(resp.Body).Decode(&measurements)
	if err != nil {
		return measurements, fmt.Errorf("failed to decode measurements: %w", err)
	}

	return measurements, nil
}
