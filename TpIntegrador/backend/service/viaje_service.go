package service

import (
	viajeClient "TPINTEGRADOR/backend/client/viaje"

	"TPINTEGRADOR/backend/dto"
	"TPINTEGRADOR/backend/model"
	e "TPINTEGRADOR/backend/utils/errors"
)

type viajeService struct{}

type viajeServiceInterface interface {
	GetViajes() (dto.ViajesDto, e.ApiError)
	InsertViaje(viajeDto dto.ViajeDto) (dto.ViajeDto, e.ApiError)
	GetViajeById(id int) (dto.ViajeDto, e.ApiError)
	GetViajesByDestino(destino string) (dto.ViajesDto, e.ApiError)
	GetViajesByNombre(nombre string) (dto.ViajesDto, e.ApiError)
}

var (
	ViajeService viajeServiceInterface
)

func init() {
	ViajeService = &viajeService{}
}

func (s *viajeService) GetViajes() (dto.ViajesDto, e.ApiError) {
	var viajes model.Viajes = viajeClient.GetViajes()
	var viajesDto dto.ViajesDto

	for _, viaje := range viajes {
		var viajeDto dto.ViajeDto
		viajeDto.ID = viaje.ID
		viajeDto.Nombre = viaje.Nombre
		viajeDto.Descripcion = viaje.Descripcion
		viajeDto.Fecha = viaje.Fecha
		viajeDto.Descripcion = viaje.Descripcion
		viajeDto.Puntaje = viaje.Puntaje
		viajeDto.Comentario = viaje.Puntaje

		viajesDto = append(viajesDto, viajeDto)
	}
	return viajesDto, nil
}

func (s *viajeService) InsertViaje(viajeDto dto.ViajeDto) (dto.ViajeDto, e.ApiError) {

	var viaje model.Viaje

	viaje.Nombre = viajeDto.Nombre
	viaje.Descripcion = viajeDto.Descripcion
	viaje.Destino = viajeDto.Destino
	viaje.Fecha = viajeDto.Fecha
	viaje.Cometario = viajeDto.Comentario
	viaje.Puntaje = viajeDto.Puntaje

	viaje = viajeClient.InsertViaje(viaje)

	viajeDto.ID = viaje.ID

	return viajeDto, nil
}

func (s *viajeService) GetViajeById(id int) (dto.ViajeDto, e.ApiError) {

	var viaje model.Viaje = viajeClient.GetViajeById(id)
	var viajeDto dto.ViajeDto

	if viaje.ID == 0 {
		return viajeDto, e.NewBadRequestApiError("Viaje No Encontrado")
	}

	viajeDto.ID = viaje.ID
	viajeDto.Nombre = viaje.Nombre
	viajeDto.Descripcion = viaje.Descripcion
	viajeDto.Destino = viaje.Destino
	viajeDto.Puntaje = viaje.Puntaje
	viajeDto.Comentario = viaje.Comentario
	viajeDto.Fecha = viaje.Fecha

	return viajeDto, nil
}

func (s *viajeService) GetViajesByNombre(nombre string) (dto.ViajesDto, e.ApiError) {

	var viajes model.Viajes = viajeClient.GetViajesByNombre(nombre)
	var viajesDto dto.ViajesDto
	var viajeDto dto.ViajeDto

	for _, viaje := range viajes {

		if viaje.Nombre == "" {
			return viajeDto, e.NewBadRequestApiError("Producto No Encontrado")
		}
		viajeDto.ID = viaje.ID
		viajeDto.Nombre = viaje.Nombre
		viajeDto.Descripcion = viaje.Descripcion
		viajeDto.Fecha = viaje.Fecha
		viajeDto.Destino = viaje.Destino
		viajeDto.Comentario = viaje.Comentario
		viajeDto.Puntaje = viaje.Puntaje

		viajesDto = append(viajesDto, viajeDto)

	}

	return viajesDto, nil
}

func (s *viajeService) GetViajesByDestino(destino string) (dto.ViajesDto, e.ApiError) {

	var viajes model.Viajes = viajeClient.GetViajesByDestino(destino)
	var viajesDto dto.ViajesDto
	var viajeDto dto.ViajeDto

	for _, viaje := range viajes {

		//if viaje.Tipo == "" {
		//	return productosDto, e.NewBadRequestApiError("Producto No Encontrado")
		//}
		viajeDto.ID = viaje.ID
		viajeDto.Nombre = viaje.Nombre
		viajeDto.Descripcion = viaje.Descripcion
		viajeDto.Fecha = viaje.Fecha
		viajeDto.Destino = viaje.Destino
		viajeDto.Comentario = viaje.Comentario
		viajeDto.Puntaje = viaje.Puntaje
		viajesDto = append(viajesDto, viajeDto)

	}

	return viajesDto, nil
}
