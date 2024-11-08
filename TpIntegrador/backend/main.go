package main

import (
	"TPINTEGRADOR/backend/app"
	"TPINTEGRADOR/backend/db"

	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Entro a main")
	db.StartDbEngine()
	log.Info("Paso el db start engine")
	app.StartRoute()
	log.Info("Paso el start route")
}
