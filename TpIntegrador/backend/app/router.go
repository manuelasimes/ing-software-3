package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.Default())
}

func StartRoute() {
	// Configura el modo de Gin según la variable de entorno GIN_MODE o por defecto a release
	// mode := gin.Mode()
	// if mode == gin.DebugMode {
	//	log.Info("Gin está en modo debug")
	// } else {
	//	log.Info("Gin está en modo release/producción")
	// }

	mapUrls()
	log.Info("Iniciando Servidor")
	router.Run(":8090")

}
