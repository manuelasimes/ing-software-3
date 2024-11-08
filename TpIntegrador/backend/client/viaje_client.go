package client

import (
	"TPINTEGRADOR/backend/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetViajeById(id int) model.Viaje {
	var viaje model.Viaje

	Db.Where("id = ?", id).First(&viaje)
	log.Debug("Viaje: ", viaje)

	return viaje
}

func GetViajesByNombre(nombre string) model.Viajes {
	var viajes model.Viajes

	Db.Where("nombre = ?", nombre).Find(&viajes)
	log.Debug("Viajes: ", viajes)

	return viajes
}

func GetViajesByDestino(destino string) model.Viajes {
	var viajes model.Viajes

	Db.Where("destino = ?", destino).Find(&viajes)
	log.Debug("Viajes: ", viajes)

	return viajes
}

func GetViajes() model.Viajes {
	var viajes model.Viajes

	Db.Find(&viajes)
	log.Debug("Viajes: ", viajes)

	return viajes
}

func InsertViaje(viaje model.Viaje) model.Viaje {
	result := Db.Create(&viaje)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("producto Creado: ", viaje.ID)
	return viaje
}
