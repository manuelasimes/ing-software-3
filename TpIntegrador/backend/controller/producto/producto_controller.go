package productoController

import (
	"backend/dto"
	service "backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Producto define la estructura de un producto
type Viaje struct {
	Destino     string `json:"destino"`
	Descripcion string `json:"descripcion"`
	Comentario  string `json:"comentario"`
}

func GetViajeById(c *gin.Context) {
	log.Debug("ID de viaje para cargar: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var viajeDto dto.ViajeDto

	viajeDto, err := service.ViajeService.GetViajeById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, viajeDto)
}

func GetViajesByNombre(c *gin.Context) {
	log.Debug("Viaje a cargar: " + c.Param("nombre"))

	nombre := c.Param("nombre")
	var viajesDto dto.viajesDto
	viajesDto, err := service.ViajeService.GetViajesByNombre(nombre)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, viajesDto)
}

func GetViajesByDestino(c *gin.Context) {
	log.Debug("Viaje a cargar: " + c.Param("destino"))

	destino := c.Param("destino")
	var viajesDto dto.viajesDto
	viajesDto, err := service.ViajeService.GetViajesByDestino(destino)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, viajesDto)
}

func GetViajes(c *gin.Context) {
	log.Info("entro al controller")
	var viajesDto dto.viajesDto
	viajesDto, err := service.ViajeService.GetViajes()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, viajesDto)
}

func InsertViaje(c *gin.Context) {
	var viajeDto dto.ViajeDto
	err := c.BindJSON(&viajeDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	viajeDto, er := service.ViajeService.InsertViaje(viajeDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, viajeDto)
}
