package app

import (
	"github.com/Manuela/TpIntegrador/backend/controller/viaje"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	log.Info("Entro a rutas producto")
	router.GET("/viaje/:id", viajeController.GetViajeById)
	router.GET("/viaje/name/:nombre", viajeController.GetViajesByNombre)
	router.GET("/viaje", viajeController.GetViajes)
	router.POST("/viaje", viajeController.InsertViaje)
	router.GET("/viaje/destino/:destino", viajeController.GetViajesByDestino)

	log.Info("termino el ruteo")
}
