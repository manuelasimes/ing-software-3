package db

import (
	viajeClient "TPINTEGRADOR/backend/client/viaje"

	"TPINTEGRADOR/backend/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "TPINTEGRADOR"
	DBUser := "root"
	DBPass := "Manuela10Simes"
	DBHost := "localhost"
	// ------------------------

	// Intenta abrir la conexión a la base de datos
	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("La conexión no se pudo abrir")
		log.Fatal(err)
	} else {
		log.Info("Conexión establecida")
	}

	// Asigna la conexión de base de datos a los clientes
	viajeClient.Db = db

}

func StartDbEngine() {
	if db == nil {
		log.Fatal("La base de datos no ha sido inicializada correctamente")
	}

	// Migrar los modelos a la base de datos
	db.AutoMigrate(&model.Viaje{})

	log.Info("Finalización de las tablas de la base de datos de migración")
}
