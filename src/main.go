package main

import (
	"fmt"
	"level-scale/dbmanager"
	"level-scale/logger"
	"level-scale/models"
	"level-scale/routes"
	"level-scale/settings"
	"net/http"
)

func main() {
	logger.Init()
	settings.Init()
	dbmanager.Init(settings.DbConfig)
	err := dbmanager.Db.AutoMigrate(&models.User{}, &models.Shop{}, &models.Product{}, &models.Order{}, &models.OrderItem{}, &models.Return{})
	if err != nil {
		logger.Log.Fatalw("Failed to setup models", "err", err)
	}
	r := routes.Init()
	logger.Log.Infow("Starting service", "port", settings.ServicePort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", settings.ServicePort), r)
	if err != nil {
		logger.Log.Fatalw("Failed to start service.", "err", err)
	}
}
