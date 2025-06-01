package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"level-scale/dbmanager"
	"level-scale/logger"
	"level-scale/metrics"
	"level-scale/models"
	"level-scale/routes"
	"level-scale/settings"
	"net/http"
)

const metricPort uint16 = 9091

func main() {
	logger.Init()
	settings.Init()

	// Database Setup
	dbmanager.Init(settings.DbConfig)
	err := dbmanager.Db.AutoMigrate(
		&models.User{},
		&models.Shop{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
		&models.Return{},
	)
	if err != nil {
		logger.Log.Fatalw("Failed to setup models", "err", err)
	}

	// Prometheus Metrics Setup
	metrics.Init()

	// Main Service APIs
	r := routes.Init()
	go func() {
		logger.Log.Infow("Starting service", "port", settings.ServicePort)
		err = http.ListenAndServe(fmt.Sprintf(":%d", settings.ServicePort), r)
		if err != nil {
			logger.Log.Fatalw("Failed to start service.", "err", err)
		}
	}()

	// Metric Service (mux'd)
	metricServer := http.NewServeMux()
	metricServer.Handle("/metrics", promhttp.Handler())
	logger.Log.Infow("Serving metrics", "port", metricPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", metricPort), metricServer)
	if err != nil {
		logger.Log.Fatalw("Failed to serve metrics", "err", err)
	}
}
